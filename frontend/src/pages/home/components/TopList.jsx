import {
  Box,
  Button,
  Card,
  CardContent,
  CardHeader,
  CardMedia,
  IconButton,
  Typography,
} from "@mui/material";
import Grid from "@mui/material/Grid2";

import React from "react";

export default function TopList() {
  return (
    <Box>
      <img src="/public/Line 1.png" alt="" />
      <Box sx={{ mt: 3 }}>
        <Typography variant="h6" color="Black">
          Лучшие товары
        </Typography>
        <Grid
          container
          spacing={{ xs: 2, md: 2 }}
          columns={{ xs: 4, sm: 4, md: 4 }}
        >
          {Array.from({ length: 4 }).map((_, index) => (
            <Grid item xs={1} sm={1} md={1} key={index}>
              <Card
                sx={{
                  width: "276px",
                  background: "#F5FCFF",
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
                    image={"/public/wheelchair.png"}
                    alt={"wheelchair"}
                    sx={{
                      width: "200px",
                      height: { xs: "200px", sm: "200px", md: "200px" },
                      objectFit: "cover",
                    }}
                  />
                </Box>

                <CardContent>
                  <CardHeader
                    title={
                      "Кресло-коляска облегчённая механическая MEYRA Eurochair 2.750 "
                    }
                  />
                  <Box
                    sx={{
                      display: "flex",
                      justifyContent: "space-between",
                      alignItems: "center",
                    }}
                  >
                    <Typography variant="h6" sx={{ color: "#004B8D" }}>
                      124456 руб. 
                    </Typography>
                    <Typography
                      variant="body2"
                      sx={{
                        color: "text.secondary",
                        textDecoration: "line-through",
                      }}
                    >
                      124456 руб. 
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
                        background: "#1E90FF",
                        borderRadius: "10px",
                      }}
                      variant="contained"
                    >
                      Подробнее
                    </Button>
                    <IconButton>
                      <img
                        style={{ width: "50px", height: "50px" }}
                        src="/public/basket.png"
                        alt=""
                      />
                    </IconButton>
                  </Box>
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
      </Box>
    </Box>
  );
}