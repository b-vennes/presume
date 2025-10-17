<?php

require 'vendor/autoload.php';

// reference the Dompdf namespace
use Dompdf\Dompdf;
use Dompdf\Options;

$htmlFile = getenv(name: 'CV_RENDER');

$cvContents = file_get_contents(filename: $htmlFile);

// instantiate and use the dompdf class
$options = new Options();
$options->set('isRemoteEnabled', TRUE);
$dompdf = new Dompdf(options: $options);
$dompdf->loadHtml(str: $cvContents);

// (Optional) Setup the paper size and orientation
$dompdf->setPaper(size: 'A4', orientation: 'portrait');

// Render the HTML as PDF
$dompdf->render();

$dompdf->stream(filename: 'render.pdf');
