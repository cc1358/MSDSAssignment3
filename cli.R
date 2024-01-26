n <- 100
start_time <- Sys.time()

sink("houses_output_r.txt")

for (i in 1:n) {
  houses <- read.csv(file = "houses_input.csv", header = TRUE)
  print(summary(houses))
}

sink()

# Calculate runtime
elapsed_time <- as.numeric(difftime(Sys.time(), start_time, units = "secs"))
cat("Script execution time:", elapsed_time, "seconds\n")
