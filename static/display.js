document.getElementById('asciiForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const text = document.getElementById('text').value;
    const font = document.getElementById('font').value;
    document.getElementById('text').scrollTop = 0;
    const errorMessage = document.getElementById('error-message-text');
    const errorOverlay = document.getElementById('error-overlay');
    const errorPopup = document.getElementById('error-popup');

    errorMessage.textContent = ''; // Clear any previous error
    errorOverlay.style.display = 'none'; // Hide the overlay
    errorPopup.style.display = 'none'; // Hide the popup

    if (!text) {
        showErrorPopup('Input text cannot be empty.');
        return;
    }

    // Assuming you're doing additional validation to detect unsupported characters
    if (!isValidText(text)) {
        showErrorPopup('Input contains unsupported characters.');
        return;
    }

    const response = await fetch('/generate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: new URLSearchParams({ text, font })
    });

    if (response.ok) {
        const result = await response.text();
        document.getElementById('output').textContent = result;
    } else {
        const errorText = await response.text();
        showErrorPopup(errorText);
    }
});

function isValidText(text) {
    // Add your validation logic for unsupported characters here
    // For example, check for non-ASCII characters
    const regex = /^[\x00-\x7F]*$/;
    return regex.test(text);
}

// Export Ascii function
function exportAscii() {
    const text = document.getElementById('output').textContent; // Get the generated ASCII text

    if (!text) {
        showErrorPopup('No generated ASCII text available for export.');
        return;
    }

    const blob = new Blob([text], { type: 'text/plain;charset=utf-8' });
    const link = document.createElement('a');
    link.href = URL.createObjectURL(blob);
    link.download = 'ascii_art.txt'; // Set the default download file name
    link.click();
}

// Now, attaching the "Got it!" event listener after the popup is displayed
function addErrorButtonListener() {
    document.getElementById('error-btn').addEventListener('click', () => {
        window.location.href = '/';  // Redirect to homepage
    });
}

// This function will be called whenever the error popup is shown
function showErrorPopup(message) {
    const errorMessage = document.getElementById('error-message-text');
    const errorOverlay = document.getElementById('error-overlay');
    const errorPopup = document.getElementById('error-popup');
    
    errorMessage.textContent = message;
    errorOverlay.style.display = 'block';
    errorPopup.style.display = 'block';
    
    // Attach the listener for "Got it!" button after the popup is displayed
    addErrorButtonListener();
}