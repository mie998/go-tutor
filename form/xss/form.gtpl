<html>
<head>
<title></title>
</head>
<body>
<form method="post" action="/readyForXSS">
    XSS: <input type="text" name="ready">
    <input type="submit" value="ログイン">
</form>
<form method="post" action="/notReadyForXSS">
    XSS: <input type="text" name="notReady">
    <input type="submit" value="ログイン">
</form>
</body>
</html>
