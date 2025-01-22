import { Box, Container, Grid, Paper, Typography } from "@mui/material";
import AdminProductTable from "../AdminProductTable/AdminProductTable";
import AdminCategoriesTable from "../AdminCategoriesTable/AdminCategoriesTable";
import AdminUserTable from "../AdminUserTable/AdminUserTable";
import AdminUsersDiagramm from "../AdminUsersDiagramm/AdminUsersDiagramm";

const MainContent = () => {
  return (
    <Box>
      <Box>
        <AdminCategoriesTable />
      </Box>

      <Box sx={{ mt: 5 }}>
        <AdminProductTable />
      </Box>
      <Box sx={{ mt: 5 }}>
        <AdminUsersDiagramm />
      </Box>
    </Box>
  );
};
export default MainContent;
