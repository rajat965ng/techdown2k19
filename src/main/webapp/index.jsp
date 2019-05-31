<!DOCTYPE html>
<html>
<head>
    <title>The World's Most Interesting Infographic</title>
    <script src="//code.jquery.com/jquery-1.11.3.min.js"></script>
    <script>
        $(document).ready(function() {
            $("#goButton").click(makeRequest);
        });
        function makeRequest() {
            // TODO: Make authorization request
            // Define properties
            var AUTH_ENDPOINT = "https://www.facebook.com/dialog/oauth";
            var RESPONSE_TYPE = "token";
            var CLIENT_ID = "[Client_ID]";
            var REDIRECT_URI = "https://wmiig.com/callback.html";
            var SCOPE = "public_profile user_posts";
            // Build authorization request endpoint
            var requestEndpoint = AUTH_ENDPOINT + "?" +
                "response_type=" + encodeURIComponent(RESPONSE_TYPE) + "&" +
                "client_id=" + encodeURIComponent(CLIENT_ID) + "&" +
                "redirect_uri=" + encodeURIComponent(REDIRECT_URI) + "&" +
                "scope=" + encodeURIComponent(SCOPE);
            // Send to authorization request endpoint
            window.location.href = requestEndpoint;

            //alert("Button clicked!");
        } </script>
</head>
<body>
<button id="goButton" type="button">Go!</button>
<div id="results"></div>
</body>
</html>

