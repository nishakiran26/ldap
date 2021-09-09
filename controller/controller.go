package controller

import (
	"fmt"
	models "ldapbackend/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap/v3"
)

var Entries []models.UserObject

func Auth(loginDetails models.Login, c *gin.Context) bool {

	fmt.Println("Auth function is running")
	baseDN := "DC=mkclindia, DC=Local"
	ldapServer := "10.1.70.104:389"
	filterDN := "(&(objectClass=*)(sAMAccountName={username}))"
	ldapUsername := "mkclindia\\ProNExTTest"
	ldapPassword := "Mkcl#5050#"
	ldapConnection, ldapConnectionError := ldap.Dial("tcp", ldapServer)
	if ldapConnectionError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": ldapConnectionError.Error(),
		})
		fmt.Println("ldapmdl connection Error: ", ldapConnectionError)
		return false
	}
	fmt.Println(ldapConnection)
	ldapBindError := ldapConnection.Bind(ldapUsername, ldapPassword)
	if ldapBindError != nil {
		fmt.Println("ldapmdl ldapBindError: ", ldapBindError)
		return false
	}
	result, searchError := ldapConnection.Search(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter(loginDetails.Username, filterDN),
		[]string{"dn", "sAMAccountName", "mail", "sn", "givenName", "mobile"},
		nil,
	))
	if searchError != nil {
		fmt.Println("ldapmcl searcherror: ", searchError)
		c.JSON(http.StatusNotFound, gin.H{"status": searchError})
		return false
	}
	if len(result.Entries) < 1 {
		fmt.Println("User does not exists")
		c.JSON(http.StatusNotFound, gin.H{"status": "User not found"})
		return false
	}
	if len(result.Entries) > 1 {
		fmt.Println("Multiple entries of same UserID")
		c.JSON(http.StatusNotFound, gin.H{"status": "multiple entries of same user"})

		return false
	}
	if userCredentialsBindError := ldapConnection.Bind(result.Entries[0].DN, loginDetails.Password); userCredentialsBindError != nil {
		fmt.Println("ldapmdl: Invalid Credentials")
		c.JSON(http.StatusNotFound, gin.H{"status": "Invalid Credentials"})
		return false
	}
	var u models.UserObject
	u.FirstName = result.Entries[0].GetAttributeValue("givenName")
	u.LastName = result.Entries[0].GetAttributeValue("sn")
	u.MobileNumber = result.Entries[0].GetAttributeValue("mobile")
	u.Email = result.Entries[0].GetAttributeValue("mail")
	u.Username = result.Entries[0].GetAttributeValue("sAMAccountName")
	fmt.Println(u)
	return true
}
func filter(needle string, filterDN string) string {
	res := strings.Replace(filterDN, "{username}", needle, -1)
	return res

}
