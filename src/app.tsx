import { createBrowserRouter, RouterProvider } from 'react-router';

const router = createBrowserRouter([
  {
    path: '/',
    element: <h1>Home Page</h1>,
  },
]);

export const App = () => {
  return <RouterProvider router={router} />;
};
