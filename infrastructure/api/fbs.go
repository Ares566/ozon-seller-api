package api

import (
	"database/sql"
	"encoding/json"
	"ozon-seller-api/infrastructure/logger"
	"ozon-seller-api/infrastructure/net"
)

type FBSRepo struct {
	client *net.OzonClient
	db    *sql.DB
}

func NewFBSRepository(client *net.OzonClient, dbc *sql.DB) *FBSRepo {
	if client == nil || dbc == nil {
		return nil
	}

	return &FBSRepo{client, dbc}
}

func (c *FBSRepo) GetUnfulfilledList() {

	list, err := c.client.GetUnfulfilledList()
	if err != nil {
		logger.Error(err)
		return
	}

	stmt, err := c.db.Prepare("INSERT INTO 1c_queue(posting) VALUE (:posting)")
	if err != nil {
		logger.Error(err)
		return
	}
	defer stmt.Close()

	for _, posting := range list {
		bytes, jmErr := json.Marshal(posting)
		if jmErr != nil {
			logger.Error(jmErr)
		} else {

			if _, stmtErr := stmt.Exec(string(bytes)); err != nil {
				logger.Error(stmtErr)
			}
		}

	}
}
//
//func (c *FBSRepo) OrderAccept(order *entity.OrderRequest, rawBody string) (entity.OrderResponse, error) {
//
//	var orderId string
//
//	// если дубль заказа, такое бывает, это нормально, отправляем уже зафиксированный id заказа
//	err := c.db.QueryRow("SELECT order_id FROM mod_order_yandex WHERE market_order_id=?", order.Order.ID).Scan(&orderId)
//	if err == nil && orderId != "" {
//		resp := entity.OrderResponse{
//			Order: struct {
//				Accepted bool
//				Id       string
//			}{Accepted: true, Id: orderId},
//		}
//		return resp, nil
//	}
//
//	sItems, err := json.Marshal(order.Order.Items)
//	currentTime := time.Now()
//	orderId = "market-fbs-" + currentTime.Format("20060102") + "-" + strconv.Itoa(order.Order.ID)
//
//	_, err = c.db.Exec("INSERT INTO mod_order_yandex (market_order_id,order_id,items,date_request,json_request,status,regionid) VALUES (?,?,?,NOW(),?,1,77)", order.Order.ID, orderId, sItems, rawBody)
//	if err != nil {
//		return entity.OrderResponse{}, err
//	}
//
//	resp := entity.OrderResponse{
//		Order: struct {
//			Accepted bool
//			Id       string
//		}{Accepted: true, Id: orderId},
//	}
//
//	return resp, nil
//}
//
//func (c *FBSRepo) OrderStatus(order *entity.StatusRequest, rawBody string) error {
//	/*
//	 CANCELLED — заказ отменен.
//	 DELIVERED — заказ получен покупателем.
//	 DELIVERY — заказ передан в службу доставки.
//	 PICKUP — заказ доставлен в пункт самовывоза.
//	 PROCESSING — заказ находится в обработке.
//	 PENDING — по заказу требуются дополнительные действия со стороны Маркета.
//	 UNPAID — заказ оформлен, но еще не оплачен (если выбрана оплата при оформлении).
//	*/
//	var err error
//	var orderId string
//
//	err = c.db.QueryRow("SELECT order_id FROM mod_order_yandex WHERE market_order_id=?", order.Order.ID).Scan(&orderId)
//	if err != nil && orderId == "" {
//		logger.Error(err)
//		return errors.New("Не найден заказ для номера " + strconv.Itoa(order.Order.ID))
//	}
//
//	switch order.Order.Status {
//
//	case "CANCELLED":
//		err = c.orderCanceled(orderId, order, rawBody)
//	case "PROCESSING":
//		err = c.orderProcessing(orderId, order, rawBody)
//	default:
//		err = c.orderChangeStatus(orderId, order)
//
//	}
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (c *FBSRepo) orderCanceled(orderId string, order *entity.StatusRequest, rawBody string) error {
//
//	aSubStatuses := map[string]string{
//		"DELIVERY_SERVICE_UNDELIVERED": "служба доставки не смогла доставить заказ",
//		"PROCESSING_EXPIRED":           "магазин не обработал заказ в течение семи дней",
//		"REPLACING_ORDER":              "покупатель решил заменить товар другим по собственной инициативе",
//		"RESERVATION_EXPIRED":          "покупатель не завершил оформление зарезервированного заказа в течение 10 минут",
//		"RESERVATION_FAILED":           "магазин не подтвердил, что готов принять заказ. Например, не ответил или ответил некорректно на запрос о передаче и принятии заказа POST /order/accept",
//		"SHOP_FAILED":                  "магазин не может выполнить заказ",
//		"USER_CHANGED_MIND":            "покупатель отменил заказ по собственным причинам",
//		"USER_NOT_PAID":                "покупатель не оплатил заказ (для типа оплаты PREPAID) в течение 30 минут",
//		"USER_REFUSED_DELIVERY":        "покупателя не устраивают условия доставки",
//		"USER_REFUSED_PRODUCT":         "покупателю не подошел товар",
//		"USER_REFUSED_QUALITY":         "покупателя не устраивает качество товара",
//		"USER_UNREACHABLE":             "не удалось связаться с покупателем",
//	}
//	currentTime := time.Now()
//	cancelTxt := "Отмена заказа со стороны Яндекса: " + aSubStatuses[order.Order.Substatus] + ". CartID=YandexMarket"
//
//	c.db.Exec("UPDATE mod_order_yandex  SET status = 3, date_rejected = NOW(), json_order=?, reject_reason=? WHERE market_order_id =?", rawBody, order.Order.Substatus, order.Order.ID)
//	c.db.Exec("UPDATE mod_order_order SET status = 4, ya_market_order_status=?, ya_market_order_rejectreason=?  WHERE ya_market_order_id=?", order.Order.Status, order.Order.Substatus, order.Order.ID)
//	c.db.Exec("INSERT INTO mod_order_comment(orderid,ddate,adminid,body) VALUES (?,?,'0',?)", orderId, currentTime.Format("2006-01-02 15:04"), cancelTxt)
//	c.db.Exec("UPDATE event_stream SET `status`=2, dt=NOW() WHERE title LIKE '%?%'", orderId)
//
//	return nil
//}
//
//func (c *FBSRepo) orderProcessing(orderId string, order *entity.StatusRequest, rawBody string) error {
//
//	mYandexRegions2FIAS := map[int]int{
//		213:   184123,
//		214:   94527,
//		215:   94036,
//		217:   94084,
//		219:   94608,
//		10716: 94618,
//		10719: 98996,
//		10723: 97804,
//		10725: 93842,
//		10733: 96681,
//		10734: 94376,
//		10735: 95842,
//		10738: 95880,
//		10740: 95240,
//		10742: 97724,
//		10743: 99928,
//		10745: 94201,
//		10746: 95689,
//		10747: 94123,
//		10748: 100158,
//		10750: 100411,
//		10752: 101198,
//		10754: 94525,
//		10755: 101643,
//		10756: 96443,
//		10758: 94536,
//		10761: 96294,
//		10765: 101841,
//		20523: 94092,
//		20571: 94039,
//		20728: 94071,
//		21621: 94087,
//		21622: 94037,
//	}
//	var paymentType string
//	orderNotes := "Код заказа маркета: " + strconv.Itoa(order.Order.ID) + "\n"
//	if order.Order.PaymentType != "" {
//		orderNotes += "Тип оплаты: " + order.Order.PaymentType + "\n "
//		paymentType += "Тип оплаты: " + order.Order.PaymentType + ", "
//	}
//	if order.Order.PaymentMethod != "" {
//		orderNotes += "Способ оплаты: " + order.Order.PaymentMethod + "\n "
//		paymentType += "Способ оплаты: " + order.Order.PaymentMethod
//	}
//	if order.Order.Delivery.ServiceName != "" {
//		orderNotes += "Служба доставки: " + order.Order.Delivery.ServiceName + "\n\n"
//	}
//
//	// список посылок
//	for _, shipment := range order.Order.Delivery.Shipments {
//		orderNotes += "Посылка №" + strconv.Itoa(shipment.ID) + ", дата отгрузки: " + shipment.ShipmentDate + "\n"
//		if shipment.Boxes != nil {
//			orderNotes += "Грузовое место: " + fmt.Sprintf("%v", shipment.Boxes)
//		}
//	}
//
//	if order.Order.Notes != "" {
//		orderNotes += "\nКомментарий заказчика: " + order.Order.Notes + "\n "
//	}
//
//	// берем наш внутренний город, дефолтно - Москва
//	iCitiID, ok := mYandexRegions2FIAS[order.Order.Delivery.Region.ID]
//	if !ok {
//		iCitiID = 184123
//	}
//	ptDelivery := order.Order.Delivery.Type + " " + order.Order.Delivery.ServiceName
//
//	// запускаем транзакцию
//	txn, err := c.db.Begin()
//	if err != nil {
//		return err
//	}
//	defer func() {
//		_ = txn.Rollback()
//	}()
//
//	// вставка в таблицу заказов
//	_, err = txn.Exec("INSERT INTO  mod_order_order (id,date_open,aname,email,phone,ptdelivery,ptdeliveryaddress,pupaid,q,status,ya_market_order_id,ya_market_order_status,ya_market_order_rejectreason,deliverycost,delivery_method,region_id,city_id, ga_source, ga_medium, ga_campaign) "+
//		"VALUES (?,NOW(),?,?,?,?,?,?,?,'1',?,?,?,?,?,?,?,'fbs.market.yandex.ru','cpa', 'API')",
//		orderId, "", "", "", ptDelivery, NewNullString(order.Order.Delivery.Region.Name), paymentType, orderNotes, order.Order.ID, order.Order.Status, order.Order.Substatus, 0, 2, 77, iCitiID)
//
//	if err != nil {
//		logger.Error(err)
//		tError := errors.New("Ошибка добавления заказа " + orderId)
//		logger.Error(tError)
//		return tError
//	}
//
//	// вставка в таблицу состава заказа
//	for _, item := range order.Order.Items {
//
//		// TODO может избавиться от этого?
//		// получается пише в mod_order_list и комплектацию и товар. А зачем? Лишний запрос к БД
//		iId := 0
//		err = c.db.QueryRow("SELECT itemid FROM cat_item_pack WHERE packid=?", item.OfferID).Scan(&iId)
//		if err != nil && iId == 0 {
//			logger.Error(err)
//			tError := errors.New("Не найден товар для комплектации " + item.OfferID)
//			logger.Error(tError)
//			return tError
//		}
//
//		// marketstate=4 это флаг, что данная комплектация есть в  FBS
//		_, err = txn.Exec("INSERT INTO mod_order_list (orderid,iid,packid,dopid,b_qty,price,marketstate)  VALUES  (?,?,?,'0',?,?,?)",
//			orderId, iId, item.OfferID, item.Count, item.Price, 4)
//
//		if err != nil {
//			logger.Error(err)
//			return err
//		}
//
//	}
//
//	// вставки во вспомогательные таблицы
//	currentTime := time.Now()
//	statusTxt := "Поступил новый заказ. ID=" + orderId + " CartID=YandexMarketFBS"
//	_, err = txn.Exec("INSERT INTO mod_order_comment(orderid,ddate,adminid,body) VALUES (?,?,'0',?)", orderId, currentTime.Format("2006-01-02 15:04:05"), statusTxt)
//	if err != nil {
//		logger.Error(err)
//		return err
//	}
//	_, err = txn.Exec("INSERT INTO `event_stream` (`title`,`status`,`dt`,`source`,`recipient_group`) VALUES (?,'1',?,'0','10')", statusTxt, currentTime.Format("2006-01-02 15:04:05"))
//	if err != nil {
//		logger.Error(err)
//		return err
//	}
//	_, err = txn.Exec("UPDATE mod_order_yandex  SET status = 2, date_accepted = NOW(), json_order=? WHERE market_order_id=?", rawBody, order.Order.ID)
//	if err != nil {
//		logger.Error(err)
//		return err
//	}
//
//	// Commit the transaction.
//	err = txn.Commit()
//	if err != nil {
//		logger.Error(err)
//		return err
//	}
//
//	return nil
//}
//
//func (c *FBSRepo) orderChangeStatus(orderId string, order *entity.StatusRequest) error {
//
//	currentTime := time.Now()
//	statusTxt := "Новый статус от Яндекса: " + order.Order.Status
//	if order.Order.Substatus != "" {
//		statusTxt = statusTxt + ", сабстатус " + order.Order.Substatus
//	}
//
//	c.db.Exec("UPDATE mod_order_order SET ya_market_order_status=? WHERE ya_market_order_id=?", order.Order.Status, order.Order.ID)
//	c.db.Exec("INSERT INTO mod_order_comment(orderid,ddate,adminid,body) VALUES (?,?,'0',?)", orderId, currentTime.Format("2006-01-02 15:04"), statusTxt)
//	return nil
//}
//
//func (c *FBSRepo) CacheRefresh() {
//	logger.Info("🛸 Начало обновления локального кеша")
//	res, err := http.Get("https://erp.santehnika-tut.ru/yama/fbs_feed_sc.xml")
//	if err != nil {
//		logger.Error(err)
//	}
//	defer res.Body.Close()
//
//	decoder := xml.NewDecoder(res.Body)
//
//	data := &entity.YmlCatalog{}
//	err = decoder.Decode(&data)
//	if err != nil {
//		logger.Error(err)
//	}
//
//	for _, offer := range data.Shop.Offers.Offer {
//		time.Sleep(time.Microsecond)
//		key := "KeyWithTTL_" + strconv.Itoa(offer.ShopSku)
//		if !c.cache.SetWithTTL(key, offer.Count, 1, time.Duration(3*time.Hour)) {
//			logger.Info("❗ошибка записи в кеш ️" + key)
//		}
//	}
//
//	logger.Info("👌🏽 Остатки по комплектациям успешно обновлены из файла fbs_feed_sc.xml.")
//	fmt.Println(len(data.Shop.Offers.Offer))
//
//}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
