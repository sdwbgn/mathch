$('#begin').collapse({toggle: false});
$('#collapse').collapse({toggle: false});
$('#al').collapse({toggle: false});

function disable() {
    $('#easy').addClass('disabled');
    $('#medium').addClass('disabled');
    $('#hard').addClass('disabled');
    $('#al').collapse('hide');
}

function enable() {
    $('#easy').removeClass('disabled');
    $('#medium').removeClass('disabled');
    $('#hard').removeClass('disabled');
}

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

function getgame(val) {
    $('#spin').show();
    $.post('get.php', {diff: val}, function (data, status) {

        current_exp = data;
        for (var i = 0; i < current_exp.length; i++) {
            if (current_exp[i] in files)
                $('#expr').append('<div class="el"><img src="' + files[current_exp[i]] + '" /></div>');
        }
        $('#game').append('<div id="sendans" class="px-5 mx-md-5 my-4"><div class="input-group"><input id="ans" type="text" class="form-control" placeholder="Answer" onkeydown="return checkChar(event);"><div class="input-group-append"><button id="send" class="btn btn-success" type="submit">Send</button></div></div></div><button id="abandon" class="btn btn-danger">Abandon</button>');
        $('#collapse').collapse('show');
        $('#send').on('click', function (e) {
            if ($('#ans').prop('disabled') === true) return;
            $('#ans').prop('disabled', true);
            $('#send').prop('disabled', true);
            $.post('check.php', {ans: $('#ans').val(), eq: current_exp}, function (data, status) {
                if (data == 1) {
                    cleangame();
                    $('#al').text("You solved it!").addClass("alert-success").collapse('show');

                } else {
                    $('#al').text("Try again!").addClass("alert-danger").collapse('show');
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
    });
}

function cleangame() {
    $('#game').collapse('hide');
    $('#sendans').remove();
    $('#abandon').remove();
    $('#expr').empty();
    current_exp = '';
    $('#begin').collapse('show');
    enable();
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
    $('#begin').collapse('hide');
    getgame(1);
});
$('#medium').on('click', function (e) {
    if ($('#easy').hasClass('disabled')) return;
    disable();
    $('#begin').collapse('hide');
    getgame(2);
});
$('#hard').on('click', function (e) {
    if ($('#hard').hasClass('disabled')) return;
    disable();
    $('#begin').collapse('hide');
    getgame(3);
});
