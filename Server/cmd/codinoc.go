// Codinoc Site
// 2nd Year Final Group Project by Group 35
//
// University of Plymouth, UK
// NSBM Green University, LK
//
// Group Members and Leader Roles
// ------------------------------
//
// N. L. M. Nishshanka | Group Leader       | 10749846
// M. P. M. Abeyrathne | Programming Leader | 10749852
// E. C. N. Nandasiri  | Testing Leader     | 10749855
// E. L. P. Prasandika | Planning Leader    | 10749850
// A. B. Navodya       | Technical Leader   | 10749851
//                     | Quality Leader     | 10749851

package main

import (
	"net/http"
	"github.com/gorilla/mux"

	_ROUTE_ "./route"
);

const (
	CONNECTION_HOST = "localhost";
	CONNECTION_PORT = "8080";
);

func main() {
	mux_router := mux.NewRouter();

	mux_router.Handle("/", _ROUTE_.R_HomePage).Methods("GET");
	mux_router.Handle("/home", _ROUTE_.R_HomePage).Methods("GET");
	mux_router.Handle("/contact", _ROUTE_.R_ContactPage).Methods("GET");
	mux_router.Handle("/about", _ROUTE_.R_AboutPage).Methods("GET");
	mux_router.Handle("/legal", _ROUTE_.R_LegalPage).Methods("GET");
	mux_router.Handle("/site_map", _ROUTE_.R_SiteMapPage).Methods("GET");

	mux_router.Handle("/sign_in", _ROUTE_.R_SignInPage).Methods("GET");
	mux_router.Handle("/sign_in/user", _ROUTE_.R_SignInUserPage).Methods("GET");
	mux_router.Handle("/sign_in/admin", _ROUTE_.R_SignInAdminPage).Methods("GET");

	mux_router.Handle("/sign_up", _ROUTE_.R_SignUpCreateServerPage).Methods("GET");
	mux_router.Handle("/sign_up/create_admin", _ROUTE_.R_SignUpCreateAdminPage).Methods("GET");
	mux_router.Handle("/sign_up/create_admin/successful", _ROUTE_.R_SignUpCreateAdminSuccessPage).Methods("GET");

	mux_router.Handle("/forgot_password", _ROUTE_.R_ForgotPasswordPage).Methods("GET");
	mux_router.Handle("/forgot_password/email_sent/{_email_}", _ROUTE_.R_ForgotPasswordEmailSentPage).Methods("GET", "PUT");

	mux_router.Handle("/password_reset/{_code_}", _ROUTE_.R_ForgotPasswordPage).Methods("GET", "PUT");
	mux_router.Handle("/password_reset/{_code_}/success", _ROUTE_.R_PasswordResetSuccessPage).Methods("GET", "PUT");
	mux_router.Handle("/password_reset/{_code_}/failed", _ROUTE_.R_PasswordResetFailedPage).Methods("GET", "PUT");

	mux_router.Handle("/sign_out", _ROUTE_.R_SignOutPage).Methods("GET");

	mux_router.Handle("/editor", _ROUTE_.R_IDEPage).Methods("GET");

	connection_error := http.ListenAndServe(CONNECTION_HOST + ":" + CONNECTION_PORT, mux_router);

	if connection_error != nil {
		return;
	}
}