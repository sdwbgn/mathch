<?php
if (isset($_POST['diff'])) {
    $command = escapeshellcmd('python3 generate.py ' . $_POST['diff']);
    $output = shell_exec($command);
    echo $output;
}