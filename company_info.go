package garoon

type CompanyInfo struct {
	Name      string `json:"name"`
	ZipCode   string `json:"zipCode"`
	Address   string `json:"address"`
	Route     string `json:"route"`
	RouteTime string `json:"routeTime"`
	RouteFare string `json:"routeFare"`
	Phone     string `json:"phone"`
}
