package services

// you can also add your repository
// before you call module in adapter like sql or redis, etc. but you have to process them firstly
// remember this is not business logic but more like technical thing
// like when you want to insert data to separate DB first for the SQL second to Elastic or something
// you can do in here
// for business logic you only write in main service file in this example in service_hello.go

// you don't have to name it repository you can name it with what you want the point is you and your team have to know what these files do
