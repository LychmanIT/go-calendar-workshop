package db

//GET USER

//auth := models.Auth{
//Username: "goadmin",
//Password: "gopassword",
//}
//
//code, err := service.GetUser(context.Background(), &auth)
//if err != nil {
//log.Fatalln("Error ", err)
//}
//log.Println("Success ", code)

//ALL EVENTS

//service := db.NewService(dbConn)
//filter := models.Filter{
//Key:   "id",
//Value: "c7e50e19-120f-4b84-9fa7-78a44b73bb05",
//}
//filter2 := models.Filter{
//Key:   "title",
//Value: "SomeTitle",
//}
//code, err := service.AllEvents(context.Background(), filter, filter2)

// ADD EVENT

//e := models.Event{
//ID:          uuid.New().String(),
//Title:       "New Year",
//Description: "Celebrate!!!",
//Time:        "321",
//Timezone:    "ua",
//Duration:    4,
//Notes:       notes,
//}
//code, err := service.AddEvent(context.Background(), &e)
//if err != nil {
//	log.Fatalln("Error ", err)
//}
//log.Println("Success ", code)

// SHOW EVENT

//service := db.NewService(dbConn)
//events, err := service.ShowEvent(context.Background(), "6743587d-cc73-418b-9123-0cf010f8128c")
//if err != nil {
//log.Fatalln("Error ", err)
//}
//log.Println("Success ", events)

// UPDATE EVENT

//e := models.Event{
//ID:          uuid.New().String(),
//Title:       "qwe123",
//Description: "Celebrate american new year!!!",
//Time:        "555",
//Timezone:    "us",
//Duration:    3,
//Notes:       []string{
//"N521ew Instance", "qw521e",
//},
//}
//
//code, err := service.UpdateEvent(context.Background(),"6a9d0e75-2b81-4f7c-b493-000d483383f3", &e)
//if err != nil {
//log.Fatalln("Error ", err)
//}
//log.Println("Success ", code)
