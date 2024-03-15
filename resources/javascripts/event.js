window.execCommandCopy = (textToCopy) => {
    const textArea = document.createElement("textarea");
    textArea.value = textToCopy;
    document.body.appendChild(textArea);
    
    textArea.select();
    textArea.setSelectionRange(0, 99999); // For mobile devices
    document.execCommand("copy");
    
    document.body.removeChild(textArea);
}

window.openInNewTab = (url) => {
    var win = window.open(url, '_blank');
    win.focus();
}