<?php
if (isset($_POST['eq']) && isset($_POST['ans'])) {
    $command = escapeshellcmd('python3 solve.py ' . $_POST['eq']);
    $output = shell_exec($command);
    try {
        $solutions = json_decode($output, true);
        $solved = false;
        foreach ($solutions as &$value) {
            if ($value == $_POST['ans'])
                $solved = true;
        }
        if ($solved) echo 1; else echo 0;
    } catch (Exception $e) {
        echo 0;
    }

} else echo 0;