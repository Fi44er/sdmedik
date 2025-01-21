import {
  Box,
  Button,
  Card,
  CardContent,
  CardHeader,
  CardMedia,
  IconButton,
  Pagination,
  Typography,
} from "@mui/material";
import Grid from "@mui/material/Grid2";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import useProductStore from "../../../store/productStore";
import SidebarFilter from "./SidebarFilter";
import useBascketStore from "../../../store/bascketStore";

export default function CatalogDynamicPage() {
  const { id } = useParams();
  const { fetchProducts, products } = useProductStore();
  const { addProductThisBascket } = useBascketStore();
  const [currentPage, setCurrentPage] = useState(1);
  const [filters, setFilters] = useState(null); // Состояние для хранения фильтров
  const [currentProducts, setCurrentProducts] = useState([]); // Переменная для хранения текущих продуктов
  const [quantity, setQuantity] = useState(0);
  const ProductsPerPage = 20;

  const category_id = id;

  useEffect(() => {
    fetchProducts(category_id, filters); // Передаем фильтры в fetchProducts
  }, [category_id, fetchProducts, filters]); // Добавляем filters в зависимости

  useEffect(() => {
    if (products?.data) {
      let normalizedProducts = [];
      if (!Array.isArray(products.data)) {
        normalizedProducts = [products.data]; // Приводим объект к массиву
      } else {
        normalizedProducts = products.data;
      }
      setCurrentProducts(normalizedProducts);
    }
  }, [products]);

  const indexOfLastItem = currentPage * ProductsPerPage;
  const indexOfFirstItem = indexOfLastItem - ProductsPerPage;

  const handleChangePage = (event, value) => {
    setCurrentPage(value);
  };

  const paginatedProducts = currentProducts.slice(
    indexOfFirstItem,
    indexOfLastItem
  );

  const hendleAddProductThithBascket = async (id) => {
    setQuantity(quantity + 1);
    const product_id = id;
    console.log(id, quantity);

    await addProductThisBascket(product_id, quantity);
  };

  return (
    <Box sx={{ mt: 1, mb: 5 }}>
      <Box sx={{ mb: 5 }}>
        <SidebarFilter setFilters={setFilters} />
      </Box>
      <Grid
        container
        spacing={{ xs: 2, md: 3 }}
        columns={{ xs: 4, sm: 4, md: 4 }}
      >
        {paginatedProducts.length > 0 ? (
          paginatedProducts.map((e) => (
            <Grid item key={e.id} xs={1} sm={1} md={1}>
              <Card
                sx={{
                  width: { xs: "100%", lg: "261px" },
                  background: "#F5FCFF",
                  cursor: "pointer",
                }}
                // onClick={() => {
                //   window.location.href = `/product/${e.id}`;
                // }}
              >
                <Box
                  sx={{
                    display: "flex",
                    justifyContent: "center",
                    alignItems: "center",
                  }}
                >
                  <CardMedia
                    component="img"
                    image={`http://127.0.0.1:8080/api/v1/image/${e.images[0].name}`}
                    alt={"wheelchair"}
                    sx={{
                      width: "200px",
                      height: { xs: "200px", sm: "200px", md: "200px" },
                      objectFit: "cover",
                    }}
                  />
                </Box>

                <CardContent>
                  <CardHeader title={e.name} />
                  <Typography variant="body2" color="text.secondary">
                    {e.article}
                  </Typography>
                  <Box
                    sx={{
                      display: "flex",
                      justifyContent: "space-between",
                      alignItems: "center",
                    }}
                  >
                    <Typography variant="h6" sx={{ color: "black" }}>
                      {e.price}
                    </Typography>
                    <Typography
                      variant="body2"
                      sx={{
                        color: "text.secondary",
                        textDecoration: "line-through",
                      }}
                    >
                      {e.price}
                    </Typography>
                  </Box>

                  <Box
                    sx={{
                      display: "flex",
                      justifyContent: "space-between",
                      alignItems: "center",
                      mt: "20px",
                    }}
                  >
                    <Button
                      sx={{
                        width: "157px",
                        height: "50px",
                        border: `2px solid #00B3A4`,
                        borderRadius: "20px",
                        color: "#00B3A4",
                      }}
                      variant="outlined"
                    >
                      В 1 клик
                    </Button>
                    <IconButton
                      onClick={() => {
                        hendleAddProductThithBascket(e.id);
                      }}
                    >
                      <img
                        style={{ width: "50px", height: "50px" }}
                        src="/public/basket_cards.png"
                        alt=""
                      />
                    </IconButton>
                  </Box>
                </CardContent>
              </Card>
            </Grid>
          ))
        ) : (
          <Typography>Нет данных для отображения</Typography>
        )}
      </Grid>
      {currentProducts.length > 0 && (
        <Pagination
          count={Math.ceil(currentProducts.length / ProductsPerPage)}
          page={currentPage}
          onChange={handleChangePage}
          sx={{
            mt: 4,
            mb: 4,
            display: "flex",
            justifyContent: "center",
            color: "#C152F0",
          }}
        />
      )}
    </Box>
  );
}
