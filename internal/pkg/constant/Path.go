package constant

const PATH_EXCEL = "excels/"
const PATH_IMAGE = "images/"
const PATH_PDF = "pdfs/"

// TODO: Hanya contoh. nanti langsung hapus saja
const PATH_TESTING = "testings/"
const PATH_CUSTOMER_ID_PHOTO = "identities/"

func PathImageCustomerID() string {
	return PATH_IMAGE + PATH_CUSTOMER_ID_PHOTO
}

func PathPDFTesting() string {
	return PATH_PDF + PATH_TESTING
}

func PathImageTesting() string {
	return PATH_IMAGE + PATH_TESTING
}

func PathExcelTesting() string {
	return PATH_EXCEL + PATH_TESTING
}
