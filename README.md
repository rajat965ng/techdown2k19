<h1>GraphQL Demo: BookSet App</h1>



<h3>Steps to Execute this App</h3>
<ul>
<li>docker run -d -p 3000:3000 rajat965ng/bookql</li>
<li>Open Postman</li>
<li>Compose POST request to <b>localhost:3000/library</b></li>
<li>Paste following payload in the body</li>
</ul>

<p>

<h3>Fetch</h3>
<p>
{
    "query" : "query getAllBooks { books {    id    name    author}}"
}
</p>
<h3>FetchByName</h3>
<p>
{
	"query":"query getBooksByName {  findByBookName(input:{name:\"Namesake\"}){    id    name    author}}"
}
</p>
<h3>Create</h3>
<p>
{
	"query":"mutation createBooks {  createBook(input:{  name:\"The White Tiger\",  author:\"Arvinda Adiga\" }){    id    name    author }}"
}
<br>
Remember that a mutation payload must return at least one field even if no data is really required.

</p>	


<h3>Update</h3>
<p>
{
	"query" :"mutation updateBook {  updateBook (input:{  id:\"1cc87605-a837-4801-ac99-f1c4b609fc10\", name: \"A white Tiger\", author: \"Arvind Adiga\" }){    id  name  author}}"
}
</p>


<h3>Delete</h3>
<p>
{
	"query":"mutation deleteBook {  deleteBook(input:{id:\"1cc87605-a837-4801-ac99-f1c4b609fc10\"}){    id}}"
}
<p>

</p>


<h3>UI Implementation for GraphQL executor is available at <i>localhost:3000/library</i></h3>

<hr>

<p>
<h3>Brief about GraphQL.</h3>
 
<h3>What is API ?</h3>
 
Most of the web and mobile applications face the requirement to fetch data from a database server. APIs are generated which serve as a bridge between the server and application. Basically, API structures the data for a client to load.
 
<h3>What is GraphQL ?</h3>
 
GraphQL can be really useful for such instances. GraphQL is a brand new API standard introduced and open-sourced by Facebook in 2015. It is more efficient and flexible substitute for building REST.
 
 
<h3>How GraphQL works?</h3>
 
GraphQL uses declarative data fetching approach where a client can request the precise data it wants from an API. It employs a single endpoint and responds with exact data the client asked for, as opposed to the multiple endpoints that REST exposes with fixed (at times, extra) data objects.
 
 
<h3>Features of GraphQL</h3>
 
<b>1. Get exactly what you want:</b>
                Sending a query to GraphQL API returns the exact data the client asked for. Nothing more, nothing less – thus results are quite predictable.
 
<b>2. Faster API response:</b>
                Since client application controls the data with their requests and not the server, GraphQL API turns out to be stable and faster. It can be insanely fast even for slow mobile network connections.
 
<b>3. Single request for multiple resources:</b>
                A veteran REST API requires making requests to multiple URL endpoints to fetch data from multiple table resources. While GraphQL can access references and fetch all the data that client needs in a single request.
 
<b>4. No more version errors:</b>
                Field columns can be added or removed from GraphQL API without influencing existing application queries. A client can utilize a single GraphQL evolving version for newer API features.
 
<b>5. API availability for multiple frameworks:</b>
                With diversified platforms and frameworks in the market, it can be really tricky to maintain a single API that would perfectly fit with changing requirements and frameworks. Since GraphQL is independent of such interfaces, it can be used on a variety of platforms.
 
<h3>What all Languages supported by GraphQL ?</h3>
 
JavaScript, Python, Java, Ruby, C#, Scala, Go, PHP & many more..
 
<h3>Who uses GraphQL?</h3>
The number of companies using GraphQL is increasing day by day. Some of them are mentioned below:
 
Facebook, Github,Twitter, Product Hunt, Yelp, Daily Motion, Myntra, Pinterest, Shopify, New York Times
 
<h3>What are the Misconceptions about GraphQL ?</h3>
 
1. One of the delusions floating around GraphQL is that it is a backend database engine. GraphQL is not a database at all; it is a Query Language for application APIs.
2. Since Facebook has always approached GraphQL and React JS together, developers have thought that they are inter-connected. This isn’t true. GraphQL is not limited to React or Native, it can be used with any technology or library that requires a client communication with API.
 
 
<h3>GraphQL v/s REST</h3>
 
<b>1. Schema and Type System:</b>
To define functions of API, GraphQL uses strong type system. All these types are coded using SDL (Schema Definition Language) of GraphQL. This schema is the communication link between client and server. The schema also allows a client to access the required data. The best part is that it works independent of both frontend and backend. Thus both the teams can concentrate on their own tasks without needing to interfere with one other after the schema is defined and understood by both the teams.
 
<b>2. Goodbye to Over and Under Fetching: (One of the major issues of REST)</b>
 
<b>Over Fetching Data:</b>
In simple terms, over fetching means that an API returns more information than the client intents to get. For instance, if we want to fetch a title of the blog post, result from REST API would also contain extra information like blog content, tags, author, etc. Such information is redundant for client application as it only needs post title string. This may increase overload on both server and client application.
 
<b>Under Fetching Data:</b>
On the other hand, under fetching means a specific API URL endpoint doesn’t return all the information required by a client. That is, the client has to make multiple requests to get each and every required information. Also, for relative information, the client would have to first fetch data from one endpoint and pass its result to the second endpoint as an argument. This can get really time taking and complex for long chains.
 
 
<b>3. Efficient Data Fetching:</b>
In REST API, one has to gather required information from multiple endpoints.
Eg. In order to fetch the title of a blog post and its comments, a client would usually need to first fetch a particular post and request another endpoint for comments based on post id for every post. But in GraphQL, this can be achieved with a single instance query.
 
<b>4. Drive Analytics:</b>
Since all the clients specify the exact information they require, excluding any redundant data, it makes it easy to drive analytics to backend by understanding what users really need. It also helps to remove specific fields that are not requested by any of the clients.
 
 
<b>5. Versioning:</b>
In REST APIs, version control is used to introduce new features. You may have noticed versioning like v1, v2, v3, etc. While this can be helpful, it increases redundancy and maintenance efforts. But GraphQL eliminates versioning all together. New features, fields or types can be appended to the existing GraphQL API itself. It doesn’t affect the existing query base.
</p>