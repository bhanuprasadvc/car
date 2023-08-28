// frontend/script.js
const carListContainer = document.getElementById("car-list");

fetch("/cars")
    .then(response => response.json())
    .then(data => {
        data.forEach(car => {
            const carCard = document.createElement("div");
            carCard.className = "car-card";
            carCard.innerHTML = `
                <h2>${car.Make} ${car.Model}</h2>
                <p>Year: ${car.Year}</p>
                <p>Availability: ${car.Availability ? "Available" : "Not Available"}</p>
            `;
            carListContainer.appendChild(carCard);
        });
    })
    .catch(error => console.error("Error fetching car data:", error));
