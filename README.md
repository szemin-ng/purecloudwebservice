# purecloudwebservice
Sample Go application that implements the Web Services Data Dip Connector contract (https://developer.mypurecloud.com/api/webservice-datadip/service-contracts.html).

Imports the following:
```
	"github.com/gorilla/mux"
```

## Instructions
### PureCloud Connector Configuration
Create a Web Services Data Dip Connector in PureCloud that points to this app's HTTP address and port. Any configuration changes made to this connector will only take effect if you restart the connector. Create the appropriate Actions you want to use. The sample Go application only implemented GetAccountByAccountNumber action. Make sure **Flatten metadata** is checked. This is required by Architect. 

### PureCloud Architect Configuration
Create an Architect call flow that calls the Bridge Actions you have configured. When the connector returns data to Architect, it is important to note that some data may be null, or in Architect, called NOT\_SET. It is important to check for NOT\_SET values because accessing it may cause the Architect call flow to fail and drop the call. For example, GetAccountByAccountNumber action returns an array of EmailAddresses. The connector may returned an empty list instead. To check if there is indeed data returned, you can:
```
If(Count(EmailAddresses) >= 1),
  If(IsSet(EmailAddresses[0],
    EmailAddresses[0],
    "",
  "")
)
```

### Running the Go application
Set PORT environment variable to the port to bind to, then:
```
purecloudwebservice
```