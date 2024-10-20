import React, { useState, useEffect } from 'react';
import axios from 'axios'; // Axios is used for making HTTP requests to the backend.
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom'; // React Router is used for managing routes in the app.
import 'bootstrap/dist/css/bootstrap.min.css'; // Import Bootstrap CSS for styling.

//Dosmailova Dinara
const TaskList = () => {
  // useState hooks to manage state variables.
  const [tasks, setTasks] = useState([]); // Holds the list of tasks.
  const [task, setTask] = useState({ title: '', description: '', completed: false }); // Manages the current task input fields.
  const [editMode, setEditMode] = useState(false); // Determines whether the user is in edit mode.
  const [editTaskId, setEditTaskId] = useState(null); // Stores the ID of the task being edited.
  const [isLoading, setIsLoading] = useState(true); // Tracks loading state.
  const [error, setError] = useState(null); // Stores any error message.

  // useEffect to fetch tasks when the component mounts.
  useEffect(() => {
    fetchTasks();
  }, []);

  //Dosmailova Dinara
  // Fetch tasks from the backend API.
  const fetchTasks = async () => {
    try {
      setIsLoading(true); // Set loading state to true while fetching.
      const response = await axios.get('http://localhost:8080/tasks'); // Make a GET request to retrieve tasks.
      setTasks(response.data); // Update the tasks state with the fetched data.
      setIsLoading(false); // Set loading state to false after fetching.
    } catch (error) {
      setError(error.message); // Capture any error and display it.
      setIsLoading(false); // Set loading state to false even if there's an error.
    }
  };

  // Handle changes in input fields.
  const handleInputChange = (e) => {
    const { name, value } = e.target; // Destructure name and value from the input.
    setTask({ ...task, [name]: value }); // Update the task state with new input values.
  };

  //Dosmailova Dinara
  // Handle form submission for adding or updating tasks.
  const handleSubmit = async (e) => {
    e.preventDefault(); // Prevent page reload on form submission.
    
    // Simple validation to ensure the task title is not empty.
    if (!task.title) {
      alert('Please enter a task title.');
      return;
    }

    try {
      // If in edit mode, update the task, otherwise create a new one.
      if (editMode) {
        await axios.put(`http://localhost:8080/tasks/${editTaskId}`, task); // PUT request to update the task.
      } else {
        await axios.post('http://localhost:8080/tasks', task); // POST request to add a new task.
      }
      
      //Dosmailova Dinara
      // Reset form fields after submission.
      setTask({ title: '', description: '', completed: false });
      setEditMode(false); // Exit edit mode after updating.
      setEditTaskId(null); // Clear the edit task ID.
      fetchTasks(); // Fetch tasks again to reflect the latest changes.
    } catch (error) {
      setError(error.message); // Handle any error during submission.
    }
  };

  // Set the task to be edited and enter edit mode.
  const handleEdit = (task) => {
    setTask(task); // Set the task to the task being edited.
    setEditMode(true); // Enable edit mode.
    setEditTaskId(task.id); // Store the ID of the task being edited.
  };

  //Dosmailova Dinara
  // Delete a task based on its ID.
  const handleDelete = async (id) => {
    try {
      await axios.delete(`http://localhost:8080/tasks/${id}`); // Send a DELETE request to the server.
      fetchTasks(); // Fetch updated tasks list after deletion.
    } catch (error) {
      setError(error.message); // Capture any error during deletion.
    }
  };

  //Dosmailova Dinara
  // Switching the completion status of a task.
  const handleToggleComplete = async (task) => {
    try {
      const updatedTask = { ...task, completed: !task.completed }; // Toggle the completed status of the task.
      await axios.put(`http://localhost:8080/tasks/${task.id}`, updatedTask); // Update the task with the new completion status.
      fetchTasks(); // Fetch the updated tasks list.
    } catch (error) {
      setError(error.message); // Handle any error during completion toggle.
    }
  };

  //Dosmailova Dinara
  // The UI of the task list component.
  return (
    <div>
      <h2>Task List</h2>
      {/* Task creation and editing form */}
      <div className="task-form">
        <form onSubmit={handleSubmit}>
          <div>
            {/* Input for task title */}
            <input
              className="form-control"
              type="text"
              name="title"
              value={task.title}
              onChange={handleInputChange}
              placeholder="Task Title"
              required // Task title is required
            />
          </div>
          <div>
            {/* Textarea for task description */}
            <textarea
              className="form-control mt-2"
              name="description"
              value={task.description}
              onChange={handleInputChange}
              placeholder="Task Description"
            />
          </div>
          {/* Button to submit form (Add or Update task) */}
          <button type="submit" className="btn btn-primary mt-2">
            {editMode ? 'Update Task' : 'Add Task'}
          </button>
        </form>
      </div>
      <hr />
      {/* Display tasks or show loading/error messages */}
      {isLoading ? (
        <p>Loading...</p> // Show loading if tasks are being fetched
      ) : error ? (
        <p>Error: {error}</p> // Show error message if there's an issue
      ) : (
        <ul className="list-unstyled">
          {/* Iterate over tasks and display them */}
          {tasks.map(task => (
            <li key={task.id} className="mb-3">
              <div>
                {/* Checkbox to toggle completion status */}
                <input
                  type="checkbox"
                  checked={task.completed} // Mark as checked if task is completed
                  onChange={() => handleToggleComplete(task)} // Toggle completion on change
                />
                {/* Display task title, description, and completion status */}
                <span className="ms-2">
                  <strong>{task.title}</strong> - {task.description} -{" "}
                  <em>{task.completed ? "Completed" : "Pending"}</em> {/* Show if task is completed or pending */}
                </span>
                {/* Button to edit the task */}
                <button
                  onClick={() => handleEdit(task)}
                  className="btn btn-info btn-sm ms-2"
                >
                  Edit
                </button>
                {/* Button to delete the task */}
                <button
                  onClick={() => handleDelete(task.id)}
                  className="btn btn-danger btn-sm ms-2"
                >
                  Delete
                </button>
              </div>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

//Dosmailova Dinara
const App = () => {
  return (
    <Router>
      <div className="container">
        <h1>Task Manager</h1>
        <Link to="/">Home</Link> {/* Navigation link to home */}
        <Routes>
          <Route path="/" element={<TaskList />} /> {/* Route to render TaskList component */}
        </Routes>
      </div>
    </Router>
  );
};

export default App; // Export the App component as default.
