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
var slidersValues = document.querySelectorAll(".sliders-values p");
var blankPoints = document.getElementById("blank-stat-point")
var blankPointsMsg = document.getElementById("blank-stat-point-msg")
inputs.forEach(function(input) {
    input.addEventListener('input', function() {
        var sum = 0;
        inputs.forEach(function(input) {
            sum += parseInt(input.value);
        });
        if (input.value > 10) {
            input.value = 10;
            slidersValues.forEach(function (slider) {
                if (slider.id === input.id + "Value") {
                    slider.innerText = 10;
                }
            })
        }
        if (sum > max) {
            input.value = parseInt(input.value) - (sum - max);
            slidersValues.forEach(function (slider) {
                if (slider.id === input.id + "Value") {
                    slider.innerText = input.value;
                }
            })
        }
        let blankPts = max - sum;
        if (blankPts < 0) {
            blankPts = 0;
        }
        blankPoints.innerText = blankPts.toString();
        blankPointsMsg.textContent = "Points left: "
    });
});