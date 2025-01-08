import { Box, Container, Grid, Paper, Typography } from "@mui/material";
import AdminProductTable from "../AdminProductTable/AdminProductTable";
import AdminCategoriesTable from "../AdminCategoriesTable/AdminCategoriesTable";

const MainContent = () => {
  return (
    <Box>
      <Box>
        <Paper sx={{ p: 2 }}>
          <Typography variant="h6" gutterBottom>
            Управление категориями
          </Typography>
          <AdminCategoriesTable />
        </Paper>
      </Box>

      <Box>
        <Paper sx={{ p: 2 }}>
          <Typography variant="h6" gutterBottom>
            Управление продуктами
          </Typography>
          <AdminProductTable />
        </Paper>
      </Box>
    </Box>
  );
};
export default MainContent;
