OAuth 2.0

OAuth 2.0 allow application to exchange information securely and reliably.

Allowing user to login the application with another account.
Example pinterest allow user to login the application with facebook account. 
This is known as Federated Identity.


Allowing a application to access the resources of another application on the behalf of the user. 
Example Adobe accessing the facebook photos on your behalf.
This is known as Delegated Authority.


What were the drawbacks of OAuth 1.0 ?
1. It was overly complex.
2. It has imprecise specification that led to insecure implementations.

These reasons led to it's poor adoption.

Caution: OAuth 2.0 is not backward compatible with OAuth 1.0.


<p>
Authorization Endpoint: 

```
https://www.facebook.com/dialog/oauth
```
</p>

<p>
Create a Subdomain by putting following line at the end of <b>/etc/hosts</b> file.
<br>

```127.0.0.1       wmiig.com```
</p>
<p>
[INFO]<br>
In tomcat, 8443 is the default port that open SSL Text Service. So, you can access service running on tomcat using <i>https://wmiig.com:8443/</i> .
<br>
The callback/respond API needs to be configured accordingly at OAuth server. But in this project, facebook is not responding on port 8443 properly.
</p> 

<p>
Following command is to generate keystore. The Common Name is where you put your subdomain.domain. 
<br>

```shell
keytool -genkey -noprompt -alias tomcat-localhost -keyalg RSA -keystore localhost-rsa.jks -keypass 123456 -storepass 123456 -dname "CN=wmiig.com, OU=Dev, O=Logicbig, L=Dallas, ST=TX, C=US"
```
</p>

<p>
Excute following command to start the application <br>

```maven
sudo mvn clean -Dmaven.tomcat.port=80 -Dmaven.tomcat.path=/ tomcat7:run
```
</p>
<p>
Implicit Grant Flow (Untrusted Clients)<br>

[Request]<br>
```$xslt
GET /authorize?
       response_type=token&
       client_id=s6BhdRkqt3&
       state=xyz&
       redirect_uri=https://wmiig.com/callback.html
 Host: https://www.facebook.com/dialog/oauth
```
<br>

[Response]<br>
```$xslt
http://wmiig.com/callback.html?#
           access_token=2YotnFZFEjr1zCsicMWpAA&
           state=xyz&
           token_type=bearer&
           expires_in=3600
```
</p>
<p>
Authorization Code Grant Flow (Trusted Clients)<br>

[Request: for getting access code] <br>
```$xslt
GET /authorize?
       response_type=code&
       client_id=s6BhdRkqt3&
       state=xyz&
       redirect_uri=https://wmiig.com/callback.html
 Host: https://www.facebook.com/dialog/oauth
 
 The response_type set to 'code' instead of 'token' in case of authorized code grant. 

```
<br>

[Response: for getting access code] <br>
```$xslt
http://wmiig.com/callback?
       code=ey6XeGlAMHBpFi2LQ9JHyT6xwLbMxYRvyZAR
```
<br>

[Request: for getting access token] <br>
We have an authorization code, but not yet an access token. We have one more call to make to "exchange" our authorization code for an access token.

```$xslt
curl -X POST \
  https://graph.facebook.com/oauth/access_token \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/x-www-form-urlencoded' \
  -H 'postman-token: e167c570-95e3-bd73-20ab-10b14fed6d52' \
  -d 'code=ey6XeGlAMHBpFi2LQ9JHyT6xwLbMxYRvyZAR&
  client_id=s6BhdRkqt3&
  scope=user_friends&
  client_secret=s6BhdRkqt3sdssd&
  grant_type=authorization_code&
  redirect_uri=https://wmiig.com/callback.html'
```

Client Secret and Client ID both are mandatory, because it OAuth does authentication as well to the OAuth Service Provider. This is called Basic Authentication. The client credentials
get encoded into basic Http Scheme. The Base-64 encoded value of the client credentials are in form of 

```[Client_ID] : [Client_Secret]```

The combination of 'client_id' and 'client_secret' has to be placed in header like this

```$xslt
// Add "Authorization" header with encoded client credentials
String clientCredentials = CLIENT_ID + ":" + CLIENT_SECRET;
String encodedClientCredentials = new String(Base64.encodeBase64(clientCredentials.getBytes()));
httpPost.setHeader("Authorization", "Basic " +encodedClientCredentials);
```

<br>
One thing to notice is that in this request there is no 'response_type' but it has 'grant_type=authorization_code' instead.<br>
The value of 'code' is the value obtained from previous response.<br>
The host has also got change to 'https://graph.facebook.com/oauth/access_token'.
 
<br>

[Response: for getting access token] <br>
```$xslt
{
         "access_token":"2YotnFZFEjr1zCsicMWpAA",
         "token_type":"bearer",
         "expires_in":3600,
         "refresh_token":"tGzv3JOkF0XG5Qx2TlKWIA",
}
```
<br> The value of 'refresh_token' is optional. It is used to refresh your access token in case it gets expired. It depends on the service provider that 'refresh_token'
may or may not get return. In this example we are using Facebook API which is not returning the 'refresh_token'.
</p>
<p>
The above acquired access_token is used to access the resource. On receiving the access_token, service provider validate 2 things:
1. Is the access token is expired or revoked.<br>
2. Is the associated scope covers the requested resource.

</p>
<p>
1st Implementation<br>
Access Resource by passing access_token in query parameter.

<br>

[Request: for accessing resource] <br>
```$xslt
https://graph.facebook.com/v2.5/me?
       fields=name&
       access_token=[ACCESS_TOKEN_VALUE]
```
<br>

[Response: from accessed resource] <br>
```$xslt
{
    "name": "Rajat Nigam",
    "id": "1234567899"
}
```
</p>
<p>
2nd Implementation <br>
In case if you wish to send access_token in the body then you can query the resource like this ajax call. (Client Side Implementation)<br>

```javascript
// Request profile data with access token
   $.ajax({  type: "POST",
     url: "https://graph.facebook.com/v2.5/me?fields=name",
     headers: {"Content-Type": "application/x-www-form-urlencoded"},
     data: {
       access_token: encodeURIComponent(accessToken),
       method: "get"
     },
     contentType: "application/x-www-form-urlencoded",
     success: function(data) {
       $("#response").html("Hello, " + data.name + "!");
     }
});
```
</p>
<p>
3rd Implementation is 'Authorization request header field'. This is the most prefered implementation in production. <br>
Server side implemeation can be done by implemention the following code in any server side programming language like Java, Node etc.<br>

```
curl -X GET \
  'https://graph.facebook.com/v2.5/me/feed?%0A%20%20%20%20%20limit=25' \
  -H 'authorization: Bearer [Access_Token]' 
```

</p>
