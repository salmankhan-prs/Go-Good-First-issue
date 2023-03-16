const express = require("express");
require("dotenv").config();
const marked = require("marked");
const pincode = require("./pincode");
const { db } = require("./db");

const app = express();

app.use(express.json());

const PORT = process.env.PORT || 4000;

//connect to db
db();

const pincodeAPI = `
## Pincode API

This API provides information about Indian pincodes. The data is retrieved from a MongoDB database and can be queried based on different parameters.

### Base URL

\`\`\`

https://pincode.cyclic.app/api/

\`\`\`

### Endpoints

### Get pincode details by pincode

Returns the details of a specific pincode.

- URL: **\`/pincode/:pincodeId\`**
- Method: **\`GET\`**
- Parameters:
    - **\`pincodeId\`** (required): The 6-digit pincode.
- Response:
    - Status code: **\`200\`**
    - Body:
        
        \`\`\`json
        
        {
          "officename": "Bethapudi B.O",
          "pincode": "522502",
          "officeType": "B.O",
          "Deliverystatus": "Delivery",
          "divisionname": "Guntur",
          "regionname": "Vijayawada",
          "circlename": "Andhra Pradesh",
          "Taluk": "Mangalagiri Ho",
          "Districtname": "Guntur",
          "statename": "ANDHRA PRADESH"
        }
        
        \`\`\`
        

### Get pincode details by area

Returns the details of pincodes within a specific area.

- URL: **\`/pincode/area/:search\`**
- Method: **\`GET\`**
- Parameters:
    - **\`search\`** (required): The name of the area to search for.
- Response:
    - Status code: **\`200\`**
    - Body:
        
        \`\`\`json
        
        [  {   "officename": "Bethapudi B.O",  
        "pincode": "522502",    
        "officeType": "B.O",   
        "Deliverystatus": "Delivery",   
        "divisionname": "Guntur",   
        "regionname": "Vijayawada",    
        "circlename": "Andhra Pradesh",   
        "Taluk": "Mangalagiri Ho",    
        "Districtname": "Guntur",    
        "statename": "ANDHRA PRADESH"  }, 
        
        {"officename": "Cherukupalle B.O",  
        "pincode": "522414",    
        "officeType": "B.O",   
        "Deliverystatus": "Delivery",    
        "divisionname": "Guntur",   
        "regionname": "Vijayawada",    
        "circlename": "Andhra Pradesh",   
        "Taluk": "Mangalagiri Ho",    
        "Districtname": "Guntur",    
        "statename": "ANDHRA PRADESH"  }
        ]
        
        \`\`\`
        

### Get pincode details by district

Returns the details of pincodes within a specific district.

- URL: **\`/pincode/district/:search\`**
- Method: **\`GET\`**
- Parameters:
    - **\`search\`** (required): The name of the district to search for.
- Response:
    - Status code: **\`200\`**
    - Body:
        
        \`\`\`json
        
        [
          {
            "


`;
//get the pincode details by pincode
app.get("/api/pincode/:pincodeId", async (req, res) => {
  const pincodeId = req.params.pincodeId;
  res.send(
    await pincode.find({ pincode: pincodeId }).select({ _id: 0, __v: 0 })
  );
});
// get pincode detail's for by area
app.get("/api/pincode/area/:search", async (req, res) => {
  const search = req.params.search;

  res.send(
    await pincode.aggregate([
      {
        $match: {
          officename: {
            $regex: search,
            $options: "i",
          },
        },
      },
      { $unset: ["_id", "__v"] },
    ])
  );
});
// get pincode detail's for by district
app.get("/api/pincode/district/:search", async (req, res) => {
  const search = req.params.search;
  res.send(
    await pincode.aggregate([
      {
        $match: {
          Districtname: {
            $regex: search,
            $options: "i",
          },
        },
      },
      { $unset: ["_id", "__v"] },
    ])
  );
});

app.listen(PORT, () => console.log(`server listening on ${PORT}`));
