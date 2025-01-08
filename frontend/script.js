// API URL
const apiUrl = "http://localhost:3000/webinars"; // Ganti dengan URL API Anda

// Get All Webinars
const fetchWebinars = async () => {
  try {
    const response = await fetch(apiUrl);
    const webinars = await response.json();

    const webinarList = document.getElementById("webinarList");
    webinarList.innerHTML = "";

    webinars.forEach((webinar) => {
      const listItem = document.createElement("li");
      listItem.textContent = `${webinar.judul} | ${webinar.tanggal} | ${webinar.waktu} | ${webinar.deskripsi}`;
      webinarList.appendChild(listItem);
    });
  } catch (error) {
    console.error("Error fetching webinars:", error);
  }
};

// Create Webinar
const createWebinar = async (event) => {
  event.preventDefault();

  const judul = document.getElementById("judul").value;
  const tanggal = document.getElementById("tanggal").value;
  const waktu = document.getElementById("waktu").value;
  const deskripsi = document.getElementById("deskripsi").value;
  const panitiaId = document.getElementById("panitia_id").value;

  const webinar = {
    judul,
    tanggal,
    waktu,
    deskripsi,
    panitia_id: panitiaId,
  };

  try {
    const response = await fetch(apiUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(webinar),
    });

    if (response.ok) {
      fetchWebinars(); // Refresh webinar list after creation
      alert("Webinar created successfully!");
    } else {
      alert("Failed to create webinar");
    }
  } catch (error) {
    console.error("Error creating webinar:", error);
  }
};

// Add Event Listener for Form Submission
const createWebinarForm = document.getElementById("createWebinarForm");
createWebinarForm.addEventListener("submit", createWebinar);

// Initial fetch of webinars
fetchWebinars();
