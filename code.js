
//  здесь у меня класс конструктор, если пользователь захотел создать свою статью, это у него получилось без проблем
class Article {
    constructor(title, id, desc, content) {
        this.id = id;
        this.title = title;
        this.desc = desc;
        this.content = content;
    }
}


let root = document.querySelector("#root");

function draw(obj) {
    root.insertAdjacentHTML("beforebegin", `
    <div class="article" id= ${obj.id}>
    <h1 class="article__h1"> ${obj.Title}</h1>
    <h3 class="article__h3"> ${obj.desc} </h3>
    <p class="article__p"> 
    ${obj.Content} 
    </p>
    </div>
    `)
};



async function f() {
    let response = await fetch('http://localhost:8801/articles');

    if (response.ok) { // если HTTP-статус в диапазоне 200-299
        // получаем тело ответа (см. про этот метод ниже)
        let json = await response.json();
        //   let text = await response.text(); // прочитать тело ответа как текст

        console.log(json);
        // let value = JSON.parse(json);
        console.log(typeof(json));
        console.log(json[0])
        // draw(json[0]);
        return json[0];
    } else {
        alert("Ошибка HTTP: " + response.status);
    }

}





// draw(f());
// console.log(r);
// draw(r);

