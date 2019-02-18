<html lang="ru">
    <head>
        <title>MAThCH</title>
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
        <link rel="stylesheet" href="m.css">
    </head>
    <body class="bg-light">
        <div class="container">
            <img class="py-3 d-block mx-auto mb-4" src="img/matchlogo.png" alt="" width="300">
            <div id="al" class="alert collapse" role="alert">
            </div>
            <div id="begin" class="text-center collapse show">
                <div class="btn-group" role="group">
                    <button id="easy" type="button" class="btn btn-secondary">Easy</button>
                    <button id="medium" type="button" class="btn btn-secondary">Medium</button>
                    <button id="hard" type="button" class="btn btn-secondary">Hard</button>
                </div>
            </div>
            <div class="text-center">
                <div id="spin" class="spinner-border " role="status" style="display: none;">
                    <span class="sr-only" >Loading...</span>
                </div>
            </div>
            <div id="collapse" class="collapse">
                <div id="game" class="px-10 mx-auto text-center">
                    <div id="expr">
                    </div>
                </div>
            </div>
        </div>
        <script src="https://code.jquery.com/jquery-3.3.1.min.js" defer="defer"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" defer="defer"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" defer="defer"></script>
        <script src="m.js" defer="defer"></script>
    </body>
</html>