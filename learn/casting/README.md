# Type Casting

I have the following structure

```
    func (db *Database) _CreateUser(ctx context.Context, doc interface{}, opts ...*options.InsertOneOptions) (u *User, err error) {
      res, err := db.Collection("users").InsertOne(ctx, doc)
      if err != nil {
        return
      }

      return &User{Id: res.InsertedID.(primitive.ObjectID), Name: doc.(*User).Name, Age: doc.(*User).Age}, err
    }
```
