
func searchHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Recieved")

	buf, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("An Error occured while reading body: %v", err)
	}
	fmt.Println("Body: \n", string(buf))

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(buf)
}