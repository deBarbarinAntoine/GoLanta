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

var max = 50;
var inputs = document.querySelectorAll('input[type="range"]');
inputs.forEach(function(input) {
    input.addEventListener('input', function() {
        var sum = 0;
        inputs.forEach(function(input) {
            sum += parseInt(input.value);
        });
        if (sum > max) {
            input.value = parseInt(input.value) - (sum - max);
        }
        input.max = max - (sum - parseInt(input.value));
    });
});