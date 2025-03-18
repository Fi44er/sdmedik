import { Box, Button } from "@mui/material";
import { useNavigate } from "react-router-dom";

const CatalogButtons = () => {
  const navigate = useNavigate();

  return (
    <Box
      sx={{
        width: "max-content",
        display: { xs: "none", sm: "none", md: "flex", lg: "flex" },
        alignItems: "center",
        gridGap: 10,
      }}
    >
      <Button
        variant="contained"
        onClick={(e) => {
          e.preventDefault();
          navigate("/catalog");
        }}
        sx={{
          background: `linear-gradient(95.61deg, #A5DED1 4.71%, #00B3A4 97.25%)`,
          fontSize: "16px",
        }}
      >
        Каталог
      </Button>
      <Button
        variant="contained"
        onClick={(e) => {
          e.preventDefault();
          navigate("/catalog/certificate");
        }}
        sx={{
          background: `linear-gradient(95.61deg, #A5DED1 4.71%, #00B3A4 97.25%)`,
          fontSize: "16px",
        }}
      >
        По электронному сертификату
      </Button>
    </Box>
  );
};

export default CatalogButtons;