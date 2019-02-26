<h1>GraphQL Demo: BookSet App</h1>



<h3>Steps to Execute this App</h3>
<li>
<ul>docker run -d -p 3000:3000 rajat965ng/bookql</ul>
<ul>Open Postman</ul>
<ul>Compose POST request to <b>localhost:3000/library</b></ul>
<ul>Paste following payload in the body</ul>
</li>

<p>

#Fetch
{
    "query" : "query getAllBooks { books {    id    name    author}}"
}

#FetchByName
{
	"query":"query getBooksByName {  findByBookName(input:{name:\"Namesake\"}){    id    name    author}}"
}

#Create
{
	"query":"mutation createBooks {  createBook(input:{  name:\"The White Tiger\",  author:\"Arvinda Adiga\" }){    id    name    author }}"
}	

Remember that a mutation payload must return at least one field even if no data is really required.

#Update
{
	"query" :"mutation updateBook {  updateBook (input:{  id:\"1cc87605-a837-4801-ac99-f1c4b609fc10\", name: \"A white Tiger\", author: \"Arvind Adiga\" }){    id  name  author}}"
}


#Delete
{
	"query":"mutation deleteBook {  deleteBook(input:{id:\"1cc87605-a837-4801-ac99-f1c4b609fc10\"}){    id}}"
}

</p>


<h3>UI Implementation for Executing GraphQL is available at <i>localhost:3000/library</i></h3>