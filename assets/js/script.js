let radioAvatar = document.querySelectorAll(".avatar > input")
let avatar = document.querySelectorAll(".avatar")

for (let i = 0; i < radioAvatar.length; i++) {
    avatar[i].addEventListener('click', ()=> {
        for (let j = 0; j < radioAvatar.length; j++) {
            radioAvatar[j].removeAttribute("checked")
            avatar[j].classList.remove("selected")
        }
        radioAvatar[i].setAttribute("checked", "checked")
        avatar[i].classList.add("selected")
    })
}