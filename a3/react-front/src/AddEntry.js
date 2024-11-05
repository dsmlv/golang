import React, { useState } from 'react'; // Import React and useState hook for managing component state
import axios from 'axios'; // Import axios for making HTTP requests

//Dosmailova Dinara
function AddEntry() { // Define a functional component called AddEntry
  // State variables to store the ID and name input values
  const [id, setId] = useState('');
  const [name, setName] = useState('');

  const handleSubmit = (e) => { // Function to handle form submission
    e.preventDefault(); // Prevent the default form submission behavior

    //Dosmailova Dinara
    // Send a POST request to the backend API to create a new entry
    axios.post('http://localhost:8080/create', { id, name })
      .then((response) => {
        alert('Entry created successfully'); // If the request is successful, display a success message
      })
      .catch((error) => {
        alert('Error creating entry'); // If there is an error, display an error message
      });
  };

  //Dosmailova Dinara
  // Render a form with input fields for ID and Name and a submit button
  return (
    <form onSubmit={handleSubmit}>
      {/* Input field for ID with a change handler to update the state */}
      <input value={id} onChange={(e) => setId(e.target.value)} placeholder="ID" />
      
      {/* Input field for Name with a change handler to update the state */}
      <input value={name} onChange={(e) => setName(e.target.value)} placeholder="Name" />
      
      {/* Button to submit the form */}
      <button type="submit">Create</button>
    </form>
  );
}

export default AddEntry; // Export the component for use in other parts of the application
