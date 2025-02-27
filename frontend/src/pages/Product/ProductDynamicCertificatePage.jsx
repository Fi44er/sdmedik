import {
  Box,
  Button,
  Container,
  IconButton,
  Typography,
  Paper,
  Divider,
  List,
  ListItem,
  CardMedia,
  Select,
  MenuItem,
} from "@mui/material";
import { ArrowBack, ArrowForward } from "@mui/icons-material";
import React, { useEffect, useState } from "react";
import { Swiper, SwiperSlide } from "swiper/react";
import "swiper/swiper-bundle.css"; // Импорт стилей Swiper
import useProductStore from "../../store/productStore";
import { useParams } from "react-router-dom";
import useBascketStore from "../../store/bascketStore";
import Regions from "../../constants/regionsData/regions";
import { urlPictures } from "../../constants/constants";
import { Helmet } from "react-helmet";

export default function ProductDynamicCertificatePage() {
  const [mainImageIndex, setMainImageIndex] = useState(0);
  const [images, setImages] = useState([]);
  const { fetchProductById, products } = useProductStore();
  const { addProductThisBascket } = useBascketStore();
  const [quantity, setQuantity] = useState(1);
  const [newRegion, setNewRegion] = useState("Выберите регион");
  const { id } = useParams();

  useEffect(() => {
    fetchProductById(id);
  }, [id]);

  useEffect(() => {
    if (products.data && products.data.images) {
      const fetchedImages = products.data.images.map(
        (image) => `${urlPictures}/${image.name}`
      );
      setImages(fetchedImages);
    }
  }, [products.data]);

  const handleNextImage = () => {
    setMainImageIndex((prevIndex) => (prevIndex + 1) % images.length);
  };

  const handlePrevImage = () => {
    setMainImageIndex(
      (prevIndex) => (prevIndex - 1 + images.length) % images.length
    );
  };

  const hendleAddProductThithBascket = async (id) => {
    await addProductThisBascket(id, quantity);
  };

  const hendleChangeRegion = (event) => {
    setNewRegion(event.target.value);
  };

  const handleGetCertificate = async () => {
    const iso = newRegion;
    fetchProductById(id, iso);
  };

  const renderFeatureValue = (value) => {
    if (value === "true") {
      return "Есть";
    } else if (value === "false") {
      return "Нет";
    } else if (value === null || value === undefined || value === "") {
      return "Нет данных";
    }
    return value;
  };

  return (
    <Container sx={{ mt: 5, mb: 5 }}>
      <Helmet>
        <title>{products.data ? products.data.name : "Загрузка..."}</title>
        <meta
          name="description"
          content={
            products.data ? products.data.description : "Описание товара"
          }
        />
        <meta
          name="keywords"
          content={
            products.data
              ? `${products.data.name}, ${products.data.article}, купить ${products.data.name}`
              : "товар, артикул"
          }
        />
        <meta
          property="og:title"
          content={products.data ? products.data.name : "Загрузка..."}
        />
        <meta
          property="og:description"
          content={
            products.data ? products.data.description : "Описание товара"
          }
        />
        <meta property="og:image" content={images[mainImageIndex]} />
        <meta
          property="og:url"
          content={`https://yourwebsite.com/products/${id}`}
        />
        <meta property="og:type" content="product" />
        <meta property="og:site_name" content="Your Website Name" />
        <meta name="twitter:card" content="summary_large_image" />
        <meta
          name="twitter:title"
          content={products.data ? products.data.name : "Загрузка..."}
        />
        <meta
          name="twitter:description"
          content={
            products.data ? products.data.description : "Описание товара"
          }
        />
        <meta name="twitter:image" content={images[mainImageIndex]} />
        <script type="application/ld+json">
          {JSON.stringify({
            "@context": "https://schema.org",
            "@type": "Product",
            name: products.data ? products.data.name : "Загрузка...",
            image: images[mainImageIndex],
            description: products.data
              ? products.data.description
              : "Описание товара",
            sku: products.data ? products.data.article : "Неизвестно",
            offers: {
              "@type": "Offer",
              url: `https://yourwebsite.com/products/${id}`,
              priceCurrency: "RUB",
              price: products.data ? products.data.price : "0",
              itemCondition: "https://schema.org/NewCondition",
              availability: "https://schema.org/InStock",
            },
          })}
        </script>
      </Helmet>
      <Paper
        sx={{
          display: "flex",
          justifyContent: "space-between",
          flexDirection: { xs: "column", sm: "column", md: "unset" },
          p: 2,
        }}
      >
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            width: { xs: "100%", sm: "100%", md: "50%" },
          }}
        >
          <Container>
            <Box
              sx={{ display: "flex", justifyContent: "center", marginTop: 2 }}
            >
              <Swiper spaceBetween={10} slidesPerView={1} navigation={false}>
                <SwiperSlide
                  style={{
                    display: "flex",
                    justifyContent: "center",
                  }}
                >
                  <CardMedia
                    component="img"
                    image={images[mainImageIndex]}
                    alt={`Product Image ${mainImageIndex + 1}`}
                    style={{
                      width: { xs: 300, sm: 350, md: 400 },
                      height: { xs: 300, sm: 350, md: 400 },
                      borderRadius: 10,
                      objectFit: "cover",
                    }}
                  />
                </SwiperSlide>
              </Swiper>
            </Box>
            <Box
              sx={{
                display: "flex",
                justifyContent: "space-between",
                marginTop: 2,
              }}
            >
              <IconButton onClick={handlePrevImage}>
                <ArrowBack />
              </IconButton>
              <IconButton onClick={handleNextImage}>
                <ArrowForward />
              </IconButton>
            </Box>
          </Container>

          <Box sx={{ display: "flex", justifyContent: "center", marginTop: 2 }}>
            {products.data &&
              products.data.images &&
              products.data.images.map((image, index) => (
                <Button
                  variant="outlined"
                  key={index}
                  onClick={() => setMainImageIndex(index)}
                  sx={{
                    border:
                      mainImageIndex === index ? "2px solid #00B3A4" : "none",
                    color: "#00B3A4",
                    margin: 1,
                  }}
                >
                  <CardMedia
                    image={`${urlPictures}/${image.name}`}
                    alt={`Thumbnail ${index + 1}`}
                    sx={{
                      width: { xs: "50px", lg: "100px" },
                      height: { xs: "50px", lg: "100px" },
                      objectFit: "cover",
                      borderRadius: "5px",
                    }}
                  />
                </Button>
              ))}
          </Box>
        </Box>

        <Box sx={{ width: { xs: "100%", sm: "100%", md: "50%" } }}>
          {products.data ? (
            <Box>
              <Typography variant="h5" sx={{ fontWeight: "bold" }}>
                {products.data.name}
              </Typography>
              <Typography variant="subtitle1" sx={{ color: "gray" }}>
                Артикул: {products.data.article}
              </Typography>
              <Box
                sx={{
                  display: "flex",
                  gridGap: 20,
                  mt: 2,
                  mb: 2,
                  flexDirection: { xs: "column" },
                }}
              >
                <Select
                  value={newRegion}
                  onChange={hendleChangeRegion}
                  sx={{ minWidth: 200 }}
                >
                  <MenuItem value="Выберите регион">
                    <em>Выберите регион</em>
                  </MenuItem>
                  {Regions.map((region) => (
                    <MenuItem key={region.value} value={region.value}>
                      {region.name}
                    </MenuItem>
                  ))}
                </Select>
                <Button
                  variant="outlined"
                  onClick={() => {
                    handleGetCertificate();
                  }}
                  sx={{ border: `2px solid #00B3A4`, color: "#00B3A4" }}
                >
                 Узнать стоимость сертификата
                </Button>
              </Box>
              <Typography variant="h5" sx={{ color: "#00B3A4" }}>
                Стоимость при оплате сертификатом:{" "}
                {products.data.certificate_price} ₽
              </Typography>
              {/* Other components */}
            </Box>
          ) : (
            <Typography>Loading...</Typography>
          )}
          <Box sx={{ marginTop: 2 }}>
            <Button
              sx={{
                height: "50px",
                border: `2px solid #00B3A4`,
                borderRadius: "20px",
                color: "#00B3A4",
              }}
              variant="outlined"
              onClick={() => {
                window.location.href = `/paymants/${products.data.id}`;
              }}
            >
              Купить по сертификату
            </Button>
            <IconButton
              onClick={() => {
                hendleAddProductThithBascket(products.data.id);
              }}
            >
              <img src="/basket_cards.png" alt="Добавить в корзину" />
            </IconButton>
          </Box>
          <Divider sx={{ marginY: 2 }} />
          <Box sx={{ marginTop: 2 }}>
            <Typography variant="h6">Характеристики:</Typography>
            <List>
              {products.data?.characteristic?.map((feature, index) => (
                <ListItem key={index}>
                  <Typography>
                    {feature.name} : {renderFeatureValue(feature.value)}
                  </Typography>
                </ListItem>
              ))}
            </List>
          </Box>
          <Divider sx={{ marginY: 2 }} />
        </Box>
      </Paper>
      <Paper sx={{ marginTop: 4, p: 2 }}>
        <Typography variant="h6" sx={{ fontWeight: "bold" }}>
          Описание товара:
        </Typography>
        <Typography>{products.data?.description}</Typography>
      </Paper>
    </Container>
  );
}
