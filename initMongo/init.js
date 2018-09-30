db.createUser(
    {
      user: "user1",
      pwd: "example",
      roles: ["readWrite"]
    }
);
db.createCollection("post");
