import {
  Box,
  Card,
  CardContent,
  CardHeader,
  CardMedia,
  Container,
  Typography,
} from "@mui/material";

import Grid from "@mui/material/Grid2";

import React from "react";

const Catalogs = [
  {
    id: 1,
    name: "Кресла-коляски",
    image: "/public/wheelchair.png",
  },
  {
    id: 2,
    name: "Кресло-стулья с санитарным оснащением",
    image: "/public/wheelchair.png",
  },
  {
    id: 3,
    name: "Ходунки-опоры",
    image: "/public/wheelchair.png",
  },
  {
    id: 4,
    name: `Противопролежневые
     матрас и подушки`,
    image: "/public/wheelchair.png",
  },
  {
    id: 5,
    name: "Средства реабилитации",
    image: "/public/wheelchair.png",
  },
  {
    id: 6,
    name: "Подгузники для взрослых",
    image: "/public/wheelchair.png",
  },
  {
    id: 7,
    name: "Подгузники для детей",
    image: "/public/wheelchair.png",
  },
  {
    id: 8,
    name: "Пеленки",
    image: "/public/wheelchair.png",
  },
  {
    id: 9,
    name: "Уходовая косметика и гигиена",
    image: "/public/wheelchair.png",
  },
  {
    id: 10,
    name: "Катетеры",
    image: "/public/wheelchair.png",
  },
  {
    id: 11,
    name: "Калоприемники",
    image: "/public/wheelchair.png",
  },
  {
    id: 12,
    name: "Уроприемники",
    image: "/public/wheelchair.png",
  },
  {
    id: 13,
    name: "Нарушение функции выделения",
    image: "/public/wheelchair.png",
  },
  {
    id: 14,
    name: "Средства ухода за стомой",
    image: "/public/wheelchair.png",
  },
  {
    id: 15,
    name: "Межсуставные жидкости и синовиальные протезы",
    image: "/public/wheelchair.png",
  },
  {
    id: 16,
    name: "Специальная одежда",
    image: "/public/wheelchair.png",
  },
  {
    id: 17,
    name: "Бандажи",
    image: "/public/wheelchair.png",
  },
  {
    id: 18,
    name: "Корсеты",
    image: "/public/wheelchair.png",
  },
  {
    id: 19,
    name: "Реклинаторы",
    image: "/public/wheelchair.png",
  },
  {
    id: 20,
    name: "Туторы и аппараты",
    image: "/public/wheelchair.png",
  },
  {
    id: 21,
    name: "Трости и костыли",
    image: "/public/wheelchair.png",
  },
  {
    id: 22,
    name: "Протезы и ортезы",
    image: "/public/wheelchair.png",
  },
  {
    id: 23,
    name: "Специальные устройства",
    image: "/public/wheelchair.png",
  },
];

export default function СategoriesPage() {
  return (
    <Box sx={{ mt: 5, mb: 5 }}>
      <Container>
        <Grid
          container
          spacing={{ xs: 2, md: 2 }}
          columns={{ xs: 4, sm: 4, md: 4 }}
        >
          {Catalogs.map((item) => (
            <Grid item xs={1} sm={1} md={1} key={item.id}>
              <Card
                sx={{
                  width: "276px",
                  background: "#F5FCFF",
                  borderRadius: "20px",
                  height: "350px",
                  textAlign: "center",
                  display: "flex",
                  flexDirection: "column",
                  justifyContent: "space-around",
                }}
                onClick={(e) => {
                  e.preventDefault();
                  window.location.href = `/products/${item.id}`;
                }}
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
                    image={`/public/wheelchair.png`}
                    alt={"wheelchair"}
                    sx={{
                      width: "200px",
                      height: {
                        xs: "200px",
                        sm: "200px",
                        md: "200px",
                        lg: "200px",
                      },
                      objectFit: "cover",
                    }}
                  />
                </Box>
                <Box
                  sx={{
                    display: "flex",
                    justifyContent: "center",
                    alignItems: "center",
                  }}
                >
                  <CardContent>
                    <Typography variant="h6">{item.name}</Typography>
                  </CardContent>
                </Box>
              </Card>
            </Grid>
          ))}
        </Grid>
      </Container>
    </Box>
  );
}
