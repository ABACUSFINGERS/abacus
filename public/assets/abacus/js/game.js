function changeFingers(value, suffixID = "") {
    if (Number.isInteger(value)) {
        value = value.toString();
    }

    let leftHandValue = parseInt(value[value.length - 2]);
    let rightHandValue = parseInt(value[value.length - 1]);

    if (isNaN(leftHandValue)) {
        clearLeftHand(suffixID);
    } else {
        setHand(leftHandValue * 10, suffixID);
    }

    if (rightHandValue === 0) {
        clearRightHand(suffixID);
    } else if (rightHandValue > 0) {
        setHand(rightHandValue, suffixID);
    } else {
        console.log("Info rightHand not changing for value : " + value);
    }
}

function clearRightHand(suffixID = "") {
    changeImage("right_hand" + suffixID, "right0");

}

function clearLeftHand(suffixID = "") {
    changeImage("left_hand" + suffixID, "left0");

}

function setHand(value, suffixID = "") {
    if (Number.isInteger(value) === false) {
        value = parseInt(value);
    }

    let elRightId = "right_hand" + suffixID;
    let elLeftId = "left_hand" + suffixID;

    switch (value) {
        case 1:
            changeImage(elRightId, "right1");
            break;
        case 2:
            changeImage(elRightId, "right2");
            break;
        case 3:
            changeImage(elRightId, "right3");
            break;
        case 4:
            changeImage(elRightId, "right4");
            break;
        case 5:
            changeImage(elRightId, "right5");
            break;
        case 6:
            changeImage(elRightId, "right6");
            break;
        case 7:
            changeImage(elRightId, "right7");
            break;
        case 8:
            changeImage(elRightId, "right8");
            break;
        case 9:
            changeImage(elRightId, "right9");
            break;
        case 10:
            changeImage(elLeftId, "left10");
            break;
        case 20:
            changeImage(elLeftId, "left20");
            break;
        case 30:
            changeImage(elLeftId, "left30");
            break;
        case 40:
            changeImage(elLeftId, "left40");
            break;
        case 50:
            changeImage(elLeftId, "left50");
            break;
        case 60:
            changeImage(elLeftId, "left60");
            break;
        case 70:
            changeImage(elLeftId, "left70");
            break;
        case 80:
            changeImage(elLeftId, "left80");
            break;
        case 90:
            changeImage(elLeftId, "left90");
            break;
        default:
            console.log('Error with ' + value + ' value.');
    }
}

function changeImage(elId, imageName) {
    let image = $("#" + elId);
    image.fadeOut('fast', function () {
        image.attr('src', '../../assets/abacus/images/hands/' + imageName + '.png');
        image.fadeIn('fast');
    });
}

function changeValue(elId, value) {
    let el = $("#" + elId);
    el.html(value);
    el.animate({fontSize: '6rem'}, "fast");
    el.animate({fontSize: '4rem'}, "fast");
}

function setSound(value) {
    $("#mp3-sound").html('<audio id="mp3-sound-audio" class="d-none" preload="auto" autoplay src="../../assets/abacus/audio/' + value + '.mp3"></audio>');
}

function playSound() {
    let audio = document.getElementById('mp3-sound-audio');
    audio.currentTime = 0;
    audio.play();
}

function changeProgressbar(stepIndex, stepsCount) {
    let percentage = ((stepIndex + 1) * 100) / (stepsCount - 1);
    let el = $("#progressbar");
    el.animate({width: percentage + "%"}, "fast");
}

function changeExerciseProgressbar(stepIndex, stepsCount, isSuccess) {
    let percentage = 100 / stepsCount;
    let el = $("#progressbar-container");
    if (isSuccess) {
        el.append('<div id="progress' + stepIndex + '" class="progress-bar bg-success" role="progressbar" aria-valuenow="30" aria-valuemin="0" aria-valuemax="100"></div>');

    } else {
        el.append('<div id="progress' + stepIndex + '" class="progress-bar bg-danger" role="progressbar" aria-valuenow="30" aria-valuemin="0" aria-valuemax="100"></div>');
    }
    $("#progress" + stepIndex).animate({width: percentage + "%"}, "fast");
}

function getRandomInt(min, max, notThisValues = []) {
    min = Math.ceil(min);
    max = Math.floor(max);

    let val = 0;
    if (notThisValues.length > 0) {
        val = notThisValues[0];
        while ($.inArray(val, notThisValues) !== -1) {
            val = Math.floor(Math.random() * (max - min + 1)) + min;
        }
    } else {
        val = Math.floor(Math.random() * (max - min + 1)) + min;
    }

    return parseInt(val);
}

function shuffleArray(array) {
    let currentIndex = array.length, temporaryValue, randomIndex;

    while (0 !== currentIndex) {

        randomIndex = Math.floor(Math.random() * currentIndex);
        currentIndex -= 1;

        temporaryValue = array[currentIndex];
        array[currentIndex] = array[randomIndex];
        array[randomIndex] = temporaryValue;
    }

    return array;
}

function goToURL(url) {
    window.location.replace(url);
}

function refresh() {
    location.reload();
}

function goToNextGame(currentGameOrder) {
    let nextGameOrder = parseInt(currentGameOrder) + 1;
    window.location.replace("/jeu/?id=" + nextGameOrder);
}

function setStars(errorCounter) {
    console.log(errorCounter);
    switch (errorCounter) {
        case 3:
            $("#stars i.first").toggleClass("fa-star fa-star-o");
            $("#stars i.second").toggleClass("fa-star fa-star-o");
            $("#stars i.third").toggleClass("fa-star fa-star-o");
            break;
        case 2:
            $("#stars i.second").toggleClass("fa-star fa-star-o");
            $("#stars i.third").toggleClass("fa-star fa-star-o");
            break;
        case 1:
            $("#stars i.third").toggleClass("fa-star fa-star-o");
            break;
        case 0:
            break;
        default:
            $("#stars i.first").toggleClass("fa-star fa-star-o");
            $("#stars i.second").toggleClass("fa-star fa-star-o");
            $("#stars i.third").toggleClass("fa-star fa-star-o");
            break;
    }
}

function setEncouragement(errorCounter) {
    console.log(errorCounter);
    let text = "";
    switch (errorCounter) {
        case 3:
            text = "Presque !";
            break;
        case 2:
            text = "Pas mal !";
            break;
        case 1:
            text = "Bien joué !";
            break;
        case 0:
            text = "C'est parfait !";
            break;
        default:
            text = "Essaie encore !";

    }
    $("#encouragement").text(text);
}

function setResult(gameID, errorCounter) {
    let data = {gameID, errorCounter};
    console.log(data);
    $.ajax({
        type: "POST",
        url: "/jeu/",
        data: data,
        success: function () {
        }
    });
}

function initializeKeyboard(values) {
    let html = "";
    for (let i = 0; i < values.length; i++) {
        if (i === 0 || i === 5) {
            html += '<div class="col-2 offset-1"><button class="btn btn-outline-violet btn-lg btn-block" data-value="' + values[i].toString() + '">' + values[i].toString() + '</button></div>';
        } else {
            html += '<div class="col-2"><button class="btn btn-outline-violet btn-lg btn-block" data-value="' + values[i].toString() + '">' + values[i].toString() + '</button></div>';
        }
    }

    $("#keyboard").html(html);
}
