<?php
$file = "test.txt";
file_put_contents($file, "Hello PHP!");
echo file_get_contents($file);
?>