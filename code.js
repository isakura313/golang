class Article {
    constructor(title, desc, content) {
        this.title = title;
        this.desc = desc;
        this.content = content;
    }
}


let root = document.querySelector("#root");

function draw(obj) {
    root.insertAdjacentHTML("beforebegin", `
    <div>
    <h3> ${obj.title}</h3>
    <h5> ${obj.desc} </h5>
    <p> 
    ${obj.content} 
    </p>
    </div>
    `)

};


// let article = JSON.parse(http://localhost:8801/articles.responseText);


// (async() => {
//     console.log('before start');
  
//     const response = await fetch('http://localhost:8801/articles')
//     const myJson = await response.json();
//     console.log(JSON.stringify(myJson));
    
//     console.log('after start');
//   })();
fetch('http://localhost:8801/articles')
  .then(response => response.json())
  .then(json => console.log(json))

// var invocation = new XMLHttpRequest();
// var url = 'http://localhost:8801/articles';
   
// function callOtherDomain() {
//   if(invocation) {    
//     invocation.open('GET', url, true);
//     invocation.onreadystatechange = handler;
//     invocation.send(); 
//   }
// }

// window.onload = function(){
// draw(article);
// }
