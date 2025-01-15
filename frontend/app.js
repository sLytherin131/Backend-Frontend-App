const API_BASE_URL = "http://127.0.0.1:3000"; // Pastikan tidak ada spasi tambahan di URL

// Fields CRUD
function showFieldForm() {
    document.getElementById("fieldForm").classList.remove("d-none");
}

function hideFieldForm() {
    document.getElementById("fieldForm").classList.add("d-none");
}

function loadFields() {
    fetch(`${API_BASE_URL}/fields`)
        .then(res => res.json())
        .then(data => {
            const tbody = document.getElementById("fieldsTableBody");
            tbody.innerHTML = data.map(field => `
                <tr>
                    <td>${field.id}</td>
                    <td>${field.name}</td>
                    <td>${field.type}</td>
                    <td>${field.price_per_hour}</td>
                    <td>
                        <button class="btn btn-warning btn-sm">Edit</button>
                        <button class="btn btn-danger btn-sm" onclick="deleteField(${field.id})">Delete</button>
                    </td>
                </tr>
            `).join('');
        })
        .catch(err => console.error("Error loading fields:", err));
}

function saveField() {
    const name = document.getElementById("fieldName").value;
    const type = document.getElementById("fieldType").value;
    const price = document.getElementById("fieldPrice").value;

    // Pastikan form tidak kosong
    if (!name || !type || !price) {
        alert("Please fill all fields.");
        return;
    }

    // Tampilkan log untuk memastikan data yang dikirim
    console.log("Saving field with data:", { name, type, price_per_hour: price });

    fetch(`${API_BASE_URL}/fields`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, type, price_per_hour: price })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error("Failed to save field");
        }
        return response.json();
    })
    .then(() => {
        console.log("Field saved successfully");
        hideFieldForm();
        loadFields();
    })
    .catch(err => {
        console.error("Error saving field:", err);
        alert("Failed to save field. Please try again.");
    });
}

// Delete Field function
function deleteField(id) {
    fetch(`${API_BASE_URL}/fields/${id}`, {
        method: "DELETE",
    })
    .then(() => loadFields())
    .catch(err => console.error("Error deleting field:", err));
}

// Implement similar functions for Customers and Reservations

// Load initial data
loadFields();
