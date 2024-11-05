import React, { createContext, useState } from 'react'; // Import React, createContext, and useState from 'react'

// Create a context named DataContext to manage global data
export const DataContext = createContext();

//Dosmailova Dinara
// Define a DataProvider component to wrap parts of the application that need access to the data
export function DataProvider({ children }) {
  // State variable to store the data, initialized as an empty array
  const [data, setData] = useState([]);

  // Return the context provider, passing the data and setData function to make them available to child components
  return (
    <DataContext.Provider value={{ data, setData }}>
      {children} {/* Render any child components wrapped by DataProvider */}
    </DataContext.Provider>
  );
}
