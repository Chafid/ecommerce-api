<?php
function convert($inputString) {
    //clean all the special character
    $output = preg_replace("/[^a-zA-Z0-9, ]/", " ", $inputString);
 
    return $output;
}

echo convert("hEEEy!!!! BOB!!! hOW ARE yoU?lISTen, cAn you sEND me sample???");

php?>