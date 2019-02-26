<h1>GraphQL Demo: BookSet App</h1>



<h3>Steps to Execute this App</h3>
<li>
<ul>docker run -d -p 3000:3000 rajat965ng/bookql</ul>
<ul>Open Postman</ul>
<ul>Compose POST request to <b>localhost:3000/library</b></ul>
<ul>Paste following payload in the body</ul>
</li>

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


<h3>UI Implementation for Executing GraphQL is available at <i>localhost:3000/library</i></h3>