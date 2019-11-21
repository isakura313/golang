
let root = document.querySelector("#root");

function draw(obj) {
    root.insertAdjacentHTML("beforebegin", `
    <div class="article" id= ${obj.id}>
    <h1 class="article__h1"> ${obj.Title}</h1>
    <h3 class="article__h3"> ${obj.Desc} </h3>
    <p class="article__p"> 
    ${obj.Content} 
    </p>
    </div>
    `)
};



async function f() {
    let response = await fetch('http://localhost:8801/articles');

    if (response.ok) {

        let json = await response.json();
        //   let text = await response.text(); // прочитать тело ответа как текст

        console.log(json);
        let value = JSON.parse(json);
        console.log(typeof(value));
        console.log(value.length);
        for (let i = 1; i < value.length; i++) {
            draw(value[i]);
            
        }
        
        return json[1];
    } else {
        alert("Ошибка HTTP: " + response.status);
    }

}



f();

