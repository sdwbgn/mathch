aspeed=300;

img_cache= new Array();
current_exp = '';
allowedChars = '0123456789+=-*/'
files = {
    '0': 'img/match0.png',
    '1': 'img/match1.png',
    '2': 'img/match2.png',
    '3': 'img/match3.png',
    '4': 'img/match4.png',
    '5': 'img/match5.png',
    '6': 'img/match6.png',
    '7': 'img/match7.png',
    '8': 'img/match8.png',
    '9': 'img/match9.png',
    '+': 'img/matchplus.png',
    '-': 'img/matchsub.png',
    '/': 'img/matchdiv.png',
    '*': 'img/matchmul.png',
    '=': 'img/matcheq.png',

};
for (var i = 0; i < allowedChars.length; i++) {
    img_cache[i]=new Image();
    img_cache[i].src=files[allowedChars[i]];
}

function disable() {
    $('#easy').addClass('disabled');
    $('#medium').addClass('disabled');
    $('#hard').addClass('disabled');
    $('#al').slideUp(aspeed);
}

function enable() {
    $('#easy').removeClass('disabled');
    $('#medium').removeClass('disabled');
    $('#hard').removeClass('disabled');
}

function getgame(val) {
    $('#spin').show();
    $.post('get.php', {diff: val}, function (data, status) {

        current_exp = data;
        for (var i = 0; i < current_exp.length; i++) {
            if (current_exp[i] in files)
                $('#expr').append('<div class="el"><img src="' + files[current_exp[i]] + '" /></div>');
        }
        $('#game').append('<div id="sendans" class="px-md-5 mx-md-5 my-4"><div class="input-group"><input id="ans" type="text" class="form-control" placeholder="Answer" onkeydown="return checkChar(event);"><div class="input-group-append"><button id="send" class="btn btn-success" type="submit">Send</button></div></div></div><button id="abandon" class="btn btn-danger">Abandon</button>');
        $('#send').on('click', function (e) {
            if ($('#ans').prop('disabled') === true) return;
            $('#ans').prop('disabled', true);
            $('#send').prop('disabled', true);
            $.post('check.php', {ans: $('#ans').val(), eq: current_exp}, function (data, status) {
                if (data == 1) {
                    cleangame();

                    $('#al').removeClass("alert-danger").text("You solved it!").addClass("alert-success").slideDown('show');

                } else {
                    $('#al').removeClass("alert-success").text("Try again!").addClass("alert-danger").slideDown('show');
                    $('#ans').val("");
                    $('#ans').prop('disabled', false);
                    $('#send').prop('disabled', false);
                }

            });
        });
        $('#abandon').on('click', function (e) {
            cleangame();
        });
        $('#spin').hide();
        $('#collapse').slideDown(aspeed);
    });
}

function cleangame() {
    $('#collapse').slideUp(aspeed,function () {
        $('#sendans').remove().delay(1000);
        $('#abandon').remove().delay(1000);
        $('#expr').empty().delay(1000);
        current_exp = '';
        $('#begin').slideDown(aspeed);
        enable();
    });
}

function checkChar(event) {
    var key = event.key;
    if (event.keyCode === 13 && !$('#ans').prop('disabled') === true)
        $('#send').trigger('click');
    return allowedChars.indexOf(key) !== -1 || event.keyCode === 8 || event.keyCode === 37 || event.keyCode === 38 || event.keyCode === 39 || event.keyCode === 40;
}

$('#easy').on('click', function (e) {
    if ($('#easy').hasClass('disabled')) return;
    disable();
    $('#begin').slideUp(aspeed);
    getgame(1);
});
$('#medium').on('click', function (e) {
    if ($('#easy').hasClass('disabled')) return;
    disable();
    $('#begin').slideUp(aspeed);
    getgame(2);
});
$('#hard').on('click', function (e) {
    if ($('#hard').hasClass('disabled')) return;
    disable();
    $('#begin').slideUp(aspeed);
    getgame(3);
});
