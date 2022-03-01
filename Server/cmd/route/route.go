// Domain Routing Functions

package route

import (
	"net/http"
	"github.com/gorilla/mux"
)

// Main Pages
// ----------

var R_HomePage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page"));
	},
)

var R_ContactPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Contact Page"));
	},
)

var R_AboutPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"));
	},
)

var R_LegalPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Legal Page"));
	},
)

var R_SiteMapPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Site Map Page"));
	},
)

// Sign In Pages
// -------------

var R_SignInPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Sign In Account Selection Page"));
	},
)

var R_SignInUserPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User Sign In Page"));
	},
)

var R_SignInAdminPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server Administrator Sign In Page"));
	},
)

// Sign Up Pages
// -------------

var R_SignUpCreateServerPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Sign Up Create Server Page"));
	},
)

var R_SignUpCreateAdminPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Sign Up Create Administrator Account Page"));
	},
)
var R_SignUpCreateAdminSuccessPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Sign Up Create Administrator Account Successful Page"));
	},
)


// Forgot Password Pages
// ---------------------

var R_ForgotPasswordPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Logged Account Sign Out Page"));
	},
)

var R_ForgotPasswordEmailSentPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		mux_var := mux.Vars(r);
		get_email := mux_var["_email_"];

		w.Write([]byte("Forgot Password Main Sent to " + get_email + " Page"));
	},
)

var R_PasswordResetPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		mux_var := mux.Vars(r);
		get_code := mux_var["_code_"];

		w.Write([]byte("Password Reser " + get_code));
	},
)

var R_PasswordResetSuccessPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		mux_var := mux.Vars(r);
		get_code := mux_var["_code_"];

		w.Write([]byte("Password Reser " + get_code));
	},
)

var R_PasswordResetFailedPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		mux_var := mux.Vars(r);
		get_code := mux_var["_code_"];

		w.Write([]byte("Password Reser " + get_code));
	},
)

// Sign Out Pages
// --------------

var R_SignOutPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Logged Account Sign Out Page"));
	},
)

// IDE
// ---

var R_IDEPage = http.HandlerFunc (
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("IDE Page"));
	},
)