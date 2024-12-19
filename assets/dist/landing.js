document.addEventListener('userSignUpSuccess', function (event) {
    const alertContainer = document.getElementById('alerts-container');

    // Add class to the container
    alertContainer.className = 'mt-4 p-4';

    // Clear the content of the container
    alertContainer.innerHTML = '';


    const alert = document.createElement('div');
    alert.className = 'p-4 bg-green-100 border border-green-400 text-green-700 rounded-md shadow-md';
    alert.textContent = event.detail.value;
    alertContainer.appendChild(alert);
});

document.addEventListener('userSignUpError', function (event) {
    const alertContainer = document.getElementById('alerts-container');

    // Add class to the container
    alertContainer.className = 'mt-4 p-4';

    // Clear the content of the container
    alertContainer.innerHTML = '';

    const alert = document.createElement('div');
    alert.className = 'p-4 bg-red-100 border border-red-400 text-red-700 rounded-md shadow-md';
    alert.textContent = event.detail.value;
    alertContainer.appendChild(alert);
});