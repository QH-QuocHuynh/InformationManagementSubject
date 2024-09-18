# Information Management Subject

## Trello link: https://trello.com/w/informationmanagementsubject

# I. GIỚI THIỆU

## 1. Lý do chọn `MongoDB` làm đề tài báo cáo.

- `MongoDB`, hiện đang đứng đầu trong danh sách các cơ sở dữ liệu `NoSQL` hàng đầu, đã tiếp tục khẳng định vai trò và ảnh hưởng nổi bật trong năm 2024. Quyết định chọn `MongoDB` làm chủ đề nghiên cứu không chỉ được dựa trên những yếu tố nổi bật mà còn phản ánh sự đổi mới và ứng dụng thực tiễn của công nghệ này. Dưới đây là những lý do chính hỗ trợ cho sự lựa chọn này:

#### a. Phổ biến của `MongoDB`:

- [Stack Overflow Developer Survey 2024](https://survey.stackoverflow.co/2024/technology): `MongoDB` xếp hạng thứ 2 trong danh sách các cơ sở dữ liệu `NoSQL` phổ biến nhất, đứng sau Redis.
- **DB-Engines Ranking**: Theo bảng xếp hạng của DB-Engines, `MongoDB` thường xuyên đứng trong top 5 cơ sở dữ liệu `NoSQL` hàng đầu. Tính đến tháng 9 năm 2024, `MongoDB` đứng ở vị trí thứ 5 trong bảng xếp hạng tổng quát của các cơ sở dữ liệu.
  Nguồn: [DB-Engines Ranking](https://db-engines.com/en/ranking)

#### b. Sử dụng `MongoDB` trong doanh nghiệp:

- Theo khảo sát của [Technology | 2024 Stack Overflow Developer Survey](https://survey.stackoverflow.co/2024/technology#most-popular-technologies-database-prof), `MongoDB` được hơn 25.2% lập trình viên lựa chọn sử dụng, đặc biệt phổ biến trong các ứng dụng web và dữ liệu lớn.

- Thêm vào đó, khoảng [26.7%](https://survey.stackoverflow.co/2024/technology#most-popular-technologies-database-prof) lập trình viên hiện đang theo học `MongoDB`, cho thấy sự quan tâm đáng kể đối với công nghệ này trong cộng đồng phát triển phần mềm.

Nguồn: [Stack Overflow Developer Survey 2024](https://survey.stackoverflow.co/2024/technology)

#### c. Hiệu quả trong xử lý dữ liệu lớn:

- Theo báo cáo của `MongoDB` Inc., các công ty sử dụng `MongoDB` tăng khả năng xử lý dữ liệu lớn lên đến 70% so với việc dùng các hệ quản trị truyền thống như MySQL hay PostgreSQL.

Nguồn: [MongoDB Inc. Performance Reports](https://investors.mongodb.com/news-releases/news-release-details/mongodb-inc-announces-first-quarter-fiscal-2025-financial)

#### d. Tăng trưởng của hệ sinh thái `MongoDB`:

- Hệ sinh thái `MongoDB` đã cho thấy sự tăng trưởng mạnh mẽ đến năm 2024. Trong quý kết thúc ngày 31 tháng 7 năm 2024, `MongoDB` đã báo cáo doanh thu đạt 478,11 triệu USD, với tỷ lệ tăng trưởng hàng năm là 12,82%. Trong 12 tháng vừa qua, tổng doanh thu của `MongoDB` đạt 1,82 tỷ USD, tăng 22,37% so với năm trước. Trong năm tài chính kết thúc vào ngày 31 tháng 1 năm 2024, `MongoDB` đã đạt doanh thu hàng năm là 1,68 tỷ USD, tăng 31,07% so với cùng kỳ năm trước.

- Những con số này nhấn mạnh sự gia tăng việc sử dụng `MongoDB`, cho thấy sự mở rộng liên tục của nó trên thị trường cơ sở dữ liệu đám mây và khả năng đổi mới với các sản phẩm phù hợp với nhu cầu quản lý dữ liệu hiện đại​( [Stock Analysis](https://stockanalysis.com/stocks/mdb/revenue/)).
  Nguồn: [MongoDB Annual Report 2024](https://investors.mongodb.com/2024-Annual-Report-and-Proxy-Statement)

## 1.1. Mục tiêu của đề án

### 1.1.1 Mục tiêu cho sinh viên CNTT

#### 1.1.1.1 Hiểu rõ kiến thức cơ bản về `NoSQL`

- Đối với sinh viên CNTT, mục tiêu đầu tiên là hiểu được các khái niệm cơ bản của cơ sở dữ liệu `NoSQL`, đặc biệt là sự khác biệt giữa `NoSQL` và SQL truyền thống. Điều này giúp sinh viên chuẩn bị tốt hơn cho thị trường lao động, nơi mà nhiều công ty đang áp dụng `NoSQL` để xử lý dữ liệu phi cấu trúc và quy mô lớn.

#### 1.1.1.2 Phát triển kỹ năng thiết kế và triển khai cơ sở dữ liệu

- Sinh viên sẽ học cách thiết kế các hệ thống cơ sở dữ liệu linh hoạt và hiệu quả bằng `MongoDB`, tập trung vào các yếu tố như tối ưu hóa truy vấn và quản lý dữ liệu lớn. Đây là kỹ năng thực tiễn quan trọng khi tham gia vào các dự án phát triển phần mềm hay dữ liệu.

#### 1.1.1.3 Nâng cao khả năng làm việc với dữ liệu lớn (Big Data)

- Sinh viên sẽ có cơ hội làm quen với các hệ thống dữ liệu lớn thông qua `MongoDB`, giúp họ phát triển kỹ năng phân tích và quản lý dữ liệu khổng lồ. Đây là một yếu tố quyết định khi làm việc trong các công ty công nghệ hoặc tổ chức có nhu cầu xử lý khối lượng dữ liệu lớn.

- **_Tóm tắt:_** mục tiêu của đề án không chỉ dừng lại ở việc hiểu lý thuyết mà còn tập trung vào việc thực hành các kỹ năng thiết yếu trong quản lý và thiết kế cơ sở dữ liệu. Sinh viên sẽ học cách làm việc với Big Data và ứng dụng `MongoDB` vào các dự án thực tế.

### 1.1.2. Mục tiêu cho lập trình viên đã đi làm:

#### 1.1.2.1 Nâng cao kiến thức về hệ thống cơ sở dữ liệu hiện đại

- Lập trình viên đã có kinh nghiệm sẽ tìm hiểu về các xu hướng mới trong cơ sở dữ liệu hiện đại, đặc biệt là cách `MongoDB` hỗ trợ lưu trữ và xử lý dữ liệu phi cấu trúc một cách hiệu quả hơn. Điều này giúp họ cập nhật kiến thức và công nghệ mới nhất trong ngành.

#### 1.1.2.2 Tối ưu hóa và mở rộng ứng dụng hiện có

- `MongoDB` mang lại khả năng mở rộng linh hoạt và tối ưu hiệu suất cho các ứng dụng hiện tại. Mục tiêu này giúp lập trình viên hiểu cách áp dụng `MongoDB` để cải thiện hệ thống, tăng cường khả năng mở rộng và xử lý dữ liệu.

#### 1.1.2.3 Tìm hiểu các kỹ thuật quản lý dữ liệu phi cấu trúc

- Lập trình viên sẽ nắm vững kỹ thuật quản lý dữ liệu phi cấu trúc như JSON, một yếu tố ngày càng quan trọng trong việc phát triển các hệ thống hiện đại. Việc hiểu và quản lý dữ liệu phi cấu trúc hiệu quả giúp họ triển khai các giải pháp phức tạp hơn, đáp ứng nhu cầu thực tế của doanh nghiệp.

- **_Tóm tắt:_** các mục tiêu sẽ tập trung vào việc cập nhật kiến thức và nâng cao kỹ năng áp dụng công nghệ cơ sở dữ liệu mới vào hệ thống hiện có, nhằm tối ưu hóa và nâng cao hiệu suất của các ứng dụng. `MongoDB` cung cấp khả năng quản lý dữ liệu phi cấu trúc hiệu quả, giúp lập trình viên giải quyết các bài toán phức tạp liên quan đến dữ liệu.

## 1.2. Khi nào bạn nên sử dụng MongoDB?
- MongoDB là một lựa chọn lý tưởng trong các tình huống yêu cầu khả năng mở rộng mạnh mẽ, tính linh hoạt trong cấu trúc dữ liệu, và hiệu suất cao cho các ứng dụng thời gian thực. Việc sử dụng MongoDB trong các hệ thống lớn, phức tạp, và yêu cầu hiệu suất tức thời không chỉ giúp nâng cao trải nghiệm người dùng mà còn đảm bảo khả năng mở rộng và phát triển lâu dài của hệ thống.

Bạn nên sử dụng MongoDB trong các trường hợp sau:

#### Trường hợp 1: Hệ thống cần mở rộng ngang để quản lý dữ liệu lớn và số lượng người dùng đông đảo.

Ví dụ: 
- Khi bạn cần mở rộng quy mô hệ thống một cách hiệu quả, đặc biệt là với các ứng dụng đòi hỏi lưu trữ và xử lý dữ liệu khổng lồ.
- Trong các hệ thống yêu cầu khả năng phân tán dữ liệu trên nhiều khu vực địa lý khác nhau để tối ưu hóa hiệu suất.

#### Trường hợp 2: Dữ liệu không có cấu trúc cố định và yêu cầu sự linh hoạt trong việc lưu trữ và thay đổi mô hình dữ liệu.

Ví dụ: 

- Khi bạn phát triển ứng dụng với dữ liệu có cấu trúc linh hoạt, hoặc cấu trúc dữ liệu có thể thay đổi theo thời gian.
- Trong các dự án mà việc lưu trữ nhiều loại dữ liệu phức tạp, không đồng nhất trong cùng một cơ sở dữ liệu là cần thiết.

#### Trường hợp 3: Các ứng dụng yêu cầu truy xuất dữ liệu nhanh chóng, phản hồi theo thời gian thực, và khả năng xử lý khối lượng giao dịch đồng thời lớn.

Ví dụ: 
- Khi xây dựng các ứng dụng cần phản hồi tức thì, chẳng hạn như hệ thống theo dõi vị trí, ứng dụng trò chuyện trực tuyến, hoặc các dịch vụ yêu cầu tính tức thời cao.
- Trong các hệ thống có số lượng người dùng truy cập và cập nhật dữ liệu liên tục, đòi hỏi hiệu suất truy xuất nhanh chóng.

MongoDB, với các tính năng mạnh mẽ và linh hoạt, đang trở thành một lựa chọn hàng đầu cho các lập trình viên và doanh nghiệp muốn xây dựng các ứng dụng hiện đại, hiệu suất cao, và dễ dàng mở rộng trong tương lai.



## 2. `MongoDB` là gì? và một số khái niệm cơ bản.

### 2.1 Lịch sử hình thành của `MongoDB`

#### 2.1.1. Khái quát về `MongoDB`

- `MongoDB` là một hệ quản trị cơ sở dữ liệu `NoSQL` mã nguồn mở, được thiết kế để lưu trữ và xử lý dữ liệu phi cấu trúc và bán cấu trúc. Nó hiện là một trong những cơ sở dữ liệu `NoSQL` phổ biến nhất, được sử dụng rộng rãi trong các hệ thống phân tán và các ứng dụng yêu cầu khả năng mở rộng cao [MongoDB Official Website](https://www.mongodb.com/).

#### 2.1.2. Lịch sử phát triển

Ra đời: `MongoDB` được phát triển bởi Dwight Merriman, Eliot Horowitz, và Kevin Ryan tại công ty 10gen vào năm 2007. Phiên bản đầu tiên được phát hành dưới dạng mã nguồn mở vào năm 2009 [MongoDB Documentation](https://www.mongodb.com/docs/).

Lý do ra đời: `MongoDB` được tạo ra để giải quyết các vấn đề mà hệ quản trị cơ sở dữ liệu quan hệ (RDBMS) truyền thống gặp phải, đặc biệt là khả năng mở rộng, hiệu năng xử lý và hỗ trợ dữ liệu phi cấu trúc. Các nhà sáng lập của `MongoDB` nhận thấy rằng các hệ thống dữ liệu lớn ngày càng trở nên phức tạp và các giải pháp SQL truyền thống không thể đáp ứng được nhu cầu mở rộng một cách linh hoạt [MongoDB Wiki trên GitHub](https://github.com/mongodb/mongo/wiki).

#### 2.1.3. Các bước phát triển chính

- 2007: `MongoDB` bắt đầu phát triển tại 10gen với mục tiêu cung cấp một nền tảng dịch vụ cơ sở dữ liệu.
- 2009: Phiên bản đầu tiên của `MongoDB` được phát hành dưới dạng mã nguồn mở, tập trung vào việc cung cấp một giải pháp cơ sở dữ liệu `NoSQL` linh hoạt và có khả năng mở rộng cao [MongoDB Official Website](https://www.mongodb.com/).
- 2013: 10gen đổi tên thành `MongoDB` Inc., đánh dấu sự phát triển mạnh mẽ và tập trung hoàn toàn vào sản phẩm `MongoDB` [MongoDB Blog](https://www.mongodb.com/blog).
- Hiện tại: `MongoDB` tiếp tục phát triển với các phiên bản mới, cải tiến hiệu năng, bảo mật, và tích hợp nhiều tính năng hiện đại để đáp ứng nhu cầu của các doanh nghiệp lớn.

### 2.2 Các khái niệm cơ bản của `MongoDB`

#### 2.2.1. Cơ sở dữ liệu `NoSQL`

Khái niệm: `NoSQL` (Not Only SQL) là loại hệ quản trị cơ sở dữ liệu không sử dụng mô hình quan hệ truyền thống. `MongoDB` thuộc loại `NoSQL` và được thiết kế để xử lý dữ liệu phi cấu trúc và bán cấu trúc [MongoDB Official Website](https://www.mongodb.com/).

Đặc điểm: Cơ sở dữ liệu `NoSQL` linh hoạt hơn trong việc lưu trữ dữ liệu so với các cơ sở dữ liệu quan hệ truyền thống, cho phép lưu trữ các loại dữ liệu khác nhau mà không cần định nghĩa rõ ràng về schema.

#### 2.2.2. Document (Tài liệu)

Khái niệm: Document là đơn vị cơ bản của dữ liệu trong `MongoDB`, lưu trữ dữ liệu dưới dạng JSON hoặc BSON (Binary JSON). Mỗi document là một đối tượng dữ liệu độc lập với cấu trúc có thể khác nhau [MongoDB Documentation](https://www.mongodb.com/docs/).

Ví dụ: Một document có thể trông như sau:

```
{
  "name": "Nguyen Dien Sy dao",
  "age": 25,
  "email": "24410013@ms.uit.edu.vn"
}
```

#### 2.2.3. Collection (Tập hợp)

- Khái niệm: Collection là một tập hợp các document, tương đương với một bảng trong cơ sở dữ liệu quan hệ nhưng không có cấu trúc chặt chẽ [MongoDB Documentation](https://www.mongodb.com/docs/).

Đặc điểm: Các document trong cùng một collection có thể có cấu trúc dữ liệu khác nhau mà không cần tuân theo một schema cố định.

#### 2.2.4. Database (Cơ sở dữ liệu)

- Khái niệm: Database trong `MongoDB` là tập hợp các collections và cung cấp không gian lưu trữ cho dữ liệu. Mỗi database có thể chứa nhiều collections khác nhau [MongoDB Official Website](https://www.mongodb.com/).

#### 2.2.5. BSON (Binary JSON)

- Khái niệm: BSON là định dạng nhị phân mở rộng của JSON được sử dụng trong `MongoDB`. Nó hỗ trợ nhiều kiểu dữ liệu hơn, bao gồm cả nhị phân và số nguyên 64-bit [MongoDB Documentation](https://www.mongodb.com/docs/).

Ưu điểm: BSON giúp tăng tốc độ lưu trữ và truyền dữ liệu giữa server và ứng dụng.

#### 2.2.6. Index (Chỉ mục)

- Khái niệm: Index trong `MongoDB` giúp tăng tốc độ truy vấn dữ liệu trong các collections. Chỉ mục có thể được tạo trên một hoặc nhiều trường để cải thiện hiệu năng truy vấn [MongoDB Documentation](https://www.mongodb.com/docs/).

#### 2.2.7. Aggregation (Tổng hợp dữ liệu)

- Khái niệm: Aggregation là quá trình xử lý dữ liệu nhằm thực hiện các phép toán tổng hợp như nhóm, lọc, và sắp xếp dữ liệu. `MongoDB` sử dụng pipeline để thực hiện các bước tổng hợp này [MongoDB Documentation](https://www.mongodb.com/docs/).

#### 2.2.8. Replica Set (Tập hợp bản sao)

- Khái niệm: Replica Set là một tập hợp các bản sao của cơ sở dữ liệu để đảm bảo tính khả dụng và sao lưu dữ liệu. Replica Set bao gồm một primary node chịu trách nhiệm ghi dữ liệu và nhiều secondary nodes để sao lưu dữ liệu từ primary node [MongoDB Documentation](https://www.mongodb.com/docs/).

#### 2.2.9. Sharding (Phân chia dữ liệu)

- Khái niệm: Sharding là kỹ thuật phân chia dữ liệu trên nhiều server để cải thiện khả năng mở rộng và hiệu năng của hệ thống. `MongoDB` hỗ trợ sharding để phân phối dữ liệu trên nhiều máy chủ và duy trì hiệu suất tốt khi dữ liệu lớn [MongoDB Documentation](https://www.mongodb.com/docs/).


### 2.3. Tìm hiểu về định dạng BSON và mối liên hệ với JSON
#### 2.3.1. Khái niệm về Định dạng BSON
- BSON (Binary JSON) là một định dạng lưu trữ dữ liệu nhị phân được thiết kế đặc biệt để mở rộng và cải thiện JSON (JavaScript Object Notation). BSON cung cấp một cách hiệu quả hơn để lưu trữ và truyền tải dữ liệu cấu trúc, nhờ vào việc sử dụng định dạng nhị phân thay vì văn bản.

Các Đặc Điểm Chính của BSON:
- Nhị phân hóa (Binary Format): BSON lưu trữ dữ liệu dưới dạng nhị phân, giúp tiết kiệm không gian lưu trữ và tăng tốc độ truyền tải dữ liệu so với định dạng văn bản như JSON. Định dạng nhị phân giúp giảm chi phí phân tích dữ liệu và cải thiện hiệu suất truy xuất.

- Tính mở rộng (Extensibility): BSON hỗ trợ nhiều kiểu dữ liệu hơn JSON, bao gồm các loại dữ liệu nhị phân, ngày giờ, và các đối tượng phức tạp hơn. Điều này cho phép lưu trữ các loại dữ liệu đa dạng và phức tạp mà JSON không hỗ trợ trực tiếp.

- Hỗ trợ kích thước dữ liệu lớn (Large Data Size Support): BSON có khả năng lưu trữ các đối tượng lớn hơn và các mảng dữ liệu khổng lồ, điều này là một lợi thế trong các ứng dụng yêu cầu quản lý khối lượng dữ liệu lớn và phức tạp.

#### 2.3.2. Nguồn gốc và Lịch sử Ra đời của BSON
- BSON được phát triển bởi MongoDB Inc. để cải thiện hiệu suất lưu trữ và truyền tải dữ liệu cho hệ quản trị cơ sở dữ liệu MongoDB. MongoDB là một hệ cơ sở dữ liệu NoSQL phổ biến, được thiết kế để xử lý dữ liệu phi cấu trúc và bán cấu trúc. BSON được giới thiệu vào năm 2009 cùng với việc phát hành MongoDB phiên bản đầu tiên.

- Mục tiêu chính của việc phát triển BSON là cung cấp một định dạng dữ liệu nhị phân có thể dễ dàng tích hợp với MongoDB và các hệ thống cơ sở dữ liệu khác, đồng thời cung cấp các khả năng mở rộng và hiệu suất tốt hơn so với định dạng JSON truyền thống.

#### 2.3.3. Ví dụ về BSON
Dưới đây là một ví dụ về cách dữ liệu được lưu trữ dưới định dạng BSON và JSON:

Dữ liệu JSON
```
{
  "name": "Nguyen Dien Sỹ Dao",
  "age": 25,
  "isStudent": false,
  "courses": ["Quan ly thong tin", "He dieu hanh"],
  "address": {
    "street": "123A",
    "city": "Ho Chi Minh",
    "postalCode": "12345"
  }
}

```
**Dữ liệu BSON:**

Dữ liệu BSON của cùng một đối tượng sẽ được lưu trữ dưới dạng nhị phân, và không thể được đọc trực tiếp bằng mắt như JSON. Tuy nhiên, dưới đây là mô tả tổng quát về cấu trúc của dữ liệu BSON tương ứng:

- Type Identifier: Xác định kiểu dữ liệu của mỗi phần tử (chuỗi, số, boolean, đối tượng, mảng).
- Length: Kích thước của dữ liệu trong byte.
- Data: Nội dung dữ liệu được mã hóa thành định dạng nhị phân.
**Ví dụ BSON (Mô tả cấu trúc nhị phân)**
+ Name: "Alice" (dưới dạng chuỗi nhị phân với chỉ định kiểu dữ liệu).
+ Age: 30 (dưới dạng số nguyên 32-bit).
+ IsStudent: false (dưới dạng giá trị boolean).
+ Courses: ["Quan ly thong tin", "He dieu hanh"] (dưới dạng mảng chuỗi nhị phân).
+ Address: {
    "street": "123A",
    "city": "Ho Chi Minh",
    "postalCode": "12345"
  }
+ Street: "123A" (dưới dạng chuỗi nhị phân).
+ City: "Springfield" (dưới dạng chuỗi nhị phân).
+ PostalCode: "12345" (dưới dạng chuỗi nhị phân).


#### 2.3.4. Mối Liên hệ giữa BSON và JSON
**Tương đồng:**

- Cấu trúc Dữ liệu: Cả BSON và JSON đều sử dụng cấu trúc dữ liệu dạng key-value để lưu trữ dữ liệu. Điều này giúp cả hai định dạng đều dễ hiểu và dễ sử dụng cho các nhà phát triển.
- Đơn giản và Linh hoạt: JSON và BSON đều hỗ trợ các kiểu dữ liệu cơ bản như chuỗi, số, boolean, mảng và đối tượng.
**Khác biệt:**

- Định dạng: JSON là định dạng văn bản dễ đọc và viết cho con người, trong khi BSON là định dạng nhị phân, được tối ưu hóa cho hiệu suất lưu trữ và truyền tải.
- Kiểu dữ liệu: BSON hỗ trợ nhiều kiểu dữ liệu hơn JSON, bao gồm các loại dữ liệu nhị phân và ngày giờ, cùng với khả năng lưu trữ các đối tượng và mảng lớn hơn.
- Kích thước Dữ liệu: BSON thường có kích thước nhỏ hơn JSON khi lưu trữ dữ liệu nhờ vào định dạng nhị phân.
- Hiệu suất: BSON cải thiện hiệu suất truy xuất và lưu trữ dữ liệu, đặc biệt trong các hệ thống như MongoDB, nhờ vào các phép toán nhị phân và khả năng lưu trữ hiệu quả.


## 3. So sánh `MongoDB` với SQL
- MongoDB và SQL (Structured Query Language) đại diện cho hai loại hệ quản trị cơ sở dữ liệu khác nhau: NoSQL và quan hệ. MongoDB là một cơ sở dữ liệu NoSQL, trong khi SQL liên quan đến các hệ quản trị cơ sở dữ liệu quan hệ (RDBMS). Mặc dù cả hai đều được sử dụng để lưu trữ và quản lý dữ liệu, chúng có nhiều điểm khác biệt cơ bản về thiết kế và chức năng. Báo cáo này sẽ phân tích những điểm khác biệt chính giữa MongoDB và các hệ RDBMS, đặc biệt tập trung vào khả năng mở rộng và tính không cần schema của MongoDB.

### 3.1. Điểm khác biệt chính giữa MongoDB và SQL
#### 3.1.1 Kiến trúc và mô hình dữ liệu
**MongoDB:**

- Mô hình dữ liệu: MongoDB sử dụng mô hình lưu trữ dữ liệu dạng tài liệu (document-oriented). Dữ liệu được lưu trữ dưới dạng tài liệu BSON (Binary JSON), cho phép lưu trữ các kiểu dữ liệu phức tạp và không đồng nhất.
- Cấu trúc dữ liệu: Tài liệu trong MongoDB có thể chứa các trường khác nhau, cho phép lưu trữ dữ liệu phi cấu trúc hoặc dữ liệu có cấu trúc thay đổi.
**SQL:**

- Mô hình dữ liệu: Các hệ quản trị cơ sở dữ liệu quan hệ sử dụng mô hình dữ liệu dạng bảng (table-based). Dữ liệu được lưu trữ trong các bảng với hàng và cột.
- Cấu trúc dữ liệu: Các bảng trong SQL có cấu trúc cố định và phải tuân theo một schema được xác định trước. Mỗi hàng trong bảng phải tuân theo cấu trúc cột đã định nghĩa.
#### 3.1.2. Khả năng mở rộng và tính không cần schema
#### 3.1.2.1 Khả năng mở rộng

**MongoDB:**

- Mở rộng theo chiều ngang: MongoDB hỗ trợ mở rộng theo chiều ngang (horizontal scaling) thông qua sharding. Sharding giúp phân phối dữ liệu trên nhiều máy chủ, cải thiện hiệu suất và khả năng chịu tải của hệ thống.
- Sharding: Dữ liệu được phân phối thành các shards (phân đoạn), với mỗi shard có thể lưu trữ một phần của dữ liệu. Mongos (router) điều phối các yêu cầu từ ứng dụng đến các shard phù hợp.

**SQL:**

- Mở rộng theo chiều dọc: Các hệ RDBMS truyền thống thường sử dụng mở rộng theo chiều dọc (vertical scaling), nghĩa là nâng cấp phần cứng của máy chủ hiện tại (CPU, RAM, ổ cứng).
- Khả năng mở rộng hạn chế: Mở rộng theo chiều ngang thường phức tạp hơn trong các hệ RDBMS truyền thống, đòi hỏi phải thực hiện các thay đổi lớn về cấu trúc và quản lý dữ liệu.

#### 3.1.2.2 Tính không cần schema
**MongoDB:**

- Không cần schema: MongoDB không yêu cầu một schema cố định cho dữ liệu. Tính không cần schema cho phép các tài liệu trong cùng một tập hợp (collection) có cấu trúc khác nhau.
- Tính linh hoạt: Điều này mang lại sự linh hoạt cao hơn trong việc lưu trữ và quản lý dữ liệu, cho phép dễ dàng thay đổi cấu trúc dữ liệu mà không ảnh hưởng đến toàn bộ hệ thống.
**SQL:**

- Cần schema: Các hệ RDBMS yêu cầu phải định nghĩa schema trước khi lưu trữ dữ liệu. Schema bao gồm các bảng, cột, kiểu dữ liệu và ràng buộc.
- Cấu trúc cố định: Mọi thay đổi về cấu trúc dữ liệu yêu cầu phải thay đổi schema và có thể cần thực hiện các phép chuyển đổi phức tạp, ảnh hưởng đến hiệu suất và tính khả dụng.
### 3.1.4. Ngôn ngữ truy vấn
**MongoDB:**

- Ngôn ngữ truy vấn: MongoDB sử dụng cú pháp truy vấn riêng của mình, dựa trên JSON. Các truy vấn được thực hiện bằng cách sử dụng các phương thức và biểu thức JSON.
- Tính năng truy vấn: Mặc dù MongoDB hỗ trợ các loại truy vấn phức tạp, cú pháp và cách sử dụng có thể khác biệt so với SQL truyền thống.
**SQL:**

- Ngôn ngữ truy vấn: SQL sử dụng ngôn ngữ truy vấn cấu trúc (Structured Query Language) để thực hiện các phép toán trên cơ sở dữ liệu. SQL cung cấp một cách thức tiêu chuẩn hóa để thực hiện các truy vấn, thao tác dữ liệu và định nghĩa cấu trúc.
- Cú pháp truy vấn: SQL có cú pháp cụ thể cho các phép toán như SELECT, INSERT, UPDATE, DELETE, JOIN, GROUP BY, v.v.
### 3.1.5. Tính nhất quán và phân phối
**MongoDB:**

- Nhất quán cuối cùng: MongoDB thường sử dụng mô hình nhất quán cuối cùng (eventual consistency) trong môi trường phân phối, có nghĩa là dữ liệu có thể không đồng bộ ngay lập tức giữa các bản sao.
- Khả năng phục hồi: MongoDB cung cấp tính năng sao lưu và phục hồi để bảo vệ dữ liệu trong trường hợp sự cố.
**SQL:**

- Nhất quán ngay lập tức: Các hệ RDBMS thường đảm bảo nhất quán ngay lập tức (immediate consistency) và hỗ trợ các giao dịch ACID (Atomicity, Consistency, Isolation, Durability) để đảm bảo tính toàn vẹn dữ liệu.
- Giao dịch: SQL cung cấp cơ chế giao dịch mạnh mẽ để xử lý các phép toán phức tạp và bảo đảm tính chính xác của dữ liệu.

**Kết luận**
- MongoDB và các hệ quản trị cơ sở dữ liệu quan hệ (RDBMS) có nhiều điểm khác biệt quan trọng về cách thức lưu trữ, quản lý dữ liệu, và khả năng mở rộng. MongoDB nổi bật với khả năng mở rộng linh hoạt và tính không cần schema, cho phép ứng dụng dễ dàng xử lý khối lượng dữ liệu lớn và thay đổi cấu trúc dữ liệu. Ngược lại, các hệ RDBMS thường cung cấp tính nhất quán ngay lập tức và khả năng xử lý giao dịch mạnh mẽ nhưng có thể gặp khó khăn trong việc mở rộng và thay đổi cấu trúc dữ liệu. Việc lựa chọn giữa MongoDB và các hệ RDBMS phụ thuộc vào nhu cầu cụ thể của ứng dụng và yêu cầu về dữ liệu.


---

# II. KIẾN THỨC CƠ BẢN VỀ `MONGODB`
## 1. Nghiên cứu kiến trúc tổng quan của MongoDB

- MongoDB sử dụng một số kiến trúc cơ bản để quản lý và mở rộng dữ liệu trong các hệ thống lớn. Dưới đây là các kiến trúc phổ biến:

  - **Replica Set (Nhân bản dữ liệu):** Đây là kiến trúc quan trọng giúp đảm bảo tính khả dụng cao bằng cách tạo nhiều bản sao của dữ liệu trên các máy chủ khác nhau. Nếu primary node gặp sự cố, một secondary node có thể tự động thay thế để tiếp tục các thao tác. Điều này đảm bảo không bị gián đoạn khi xảy ra lỗi hệ thống.

    - **Ví dụ thực tế:**

      - `eBay` sử dụng `Replica Set` để đảm bảo hoạt động ổn định của cơ sở dữ liệu toàn cầu, giúp họ xử lý lượng giao dịch khổng lồ hàng ngày mà không lo ngại về việc ngừng hoạt động khi có sự cố tại một máy chủ.
  - **Sharding (Phân mảnh dữ liệu):** Kiến trúc này cho phép MongoDB phân chia dữ liệu thành nhiều phần nhỏ gọi là shards, giúp cải thiện hiệu năng và khả năng mở rộng của hệ thống. Mỗi shard sẽ được phân phối trên các máy chủ khác nhau, giúp tăng khả năng xử lý dữ liệu lớn một cách hiệu quả.

    - **Ví dụ thực tế:**

      - `Coca-Cola` sử dụng kiến trúc `Sharding` trong MongoDB để quản lý dữ liệu khổng lồ từ các chiến dịch quảng cáo và các tương tác khách hàng trực tuyến toàn cầu, đảm bảo hiệu suất cao và đáp ứng nhanh chóng ngay cả khi lượng người dùng tăng vọt.
  - **Aggregation Framework (Khung xử lý tập hợp):** Đây là công cụ mạnh mẽ trong MongoDB giúp xử lý các tập hợp dữ liệu lớn và phức tạp. `Aggregation Framework` được sử dụng rộng rãi trong các ứng dụng phân tích dữ liệu mà không cần đưa dữ liệu ra ngoài cơ sở dữ liệu.

    - **Ví dụ thực tế:**

      - Các nền tảng thương mại điện tử như `Walmart` sử dụng `Aggregation Framework` để thực hiện phân tích dữ liệu khách hàng và hàng hóa, giúp tối ưu hóa quy trình bán hàng và quản lý kho hiệu quả hơn.
**Kiến trúc phổ biến nhất:**
- Trong thực tế, Replica Set là kiến trúc được sử dụng nhiều nhất. Với khả năng đảm bảo tính sẵn sàng và khả năng khôi phục dữ liệu nhanh chóng khi gặp sự cố, nó phù hợp với các hệ thống cần tính liên tục và ổn định. Ví dụ, các công ty như MetLife đã triển khai Replica Set để đảm bảo hệ thống của họ không bị gián đoạn trong quá trình phục vụ khách hàng.

## 2. Quản lý dữ liệu của Document trong MongoDB
MongoDB quản lý dữ liệu theo mô hình Document-Oriented (hướng tài liệu), với các tài liệu được lưu trữ dưới dạng BSON (Binary JSON), cho phép lưu trữ linh hoạt, dễ dàng quản lý dữ liệu phi cấu trúc hoặc dữ liệu có cấu trúc phức tạp.

### 2.1. Cách lưu trữ dữ liệu
- Document (Tài liệu): Tài liệu là đơn vị lưu trữ cơ bản, giống như một hàng trong cơ sở dữ liệu quan hệ nhưng linh hoạt hơn, với khả năng lưu trữ các mảng, đối tượng nhúng, và các kiểu dữ liệu khác nhau. Điều này giúp MongoDB phù hợp với các loại dữ liệu không cố định.
Ví dụ: Một document có thể chứa thông tin khách hàng và đơn hàng của họ mà không cần phải tạo các bảng liên kết như trong RDBMS.
### 2.2. Quản lý Collections
- Collection (Bộ sưu tập): Các document được lưu trữ trong collections, tương đương với các bảng trong hệ thống quan hệ. Tuy nhiên, các document trong cùng một collection không cần phải có cùng cấu trúc hoặc các kiểu dữ liệu giống nhau, tạo sự linh hoạt lớn hơn.
Ví dụ: Trong một collection chứa thông tin về sản phẩm, bạn có thể lưu trữ dữ liệu cho các loại sản phẩm khác nhau mà không cần tuân theo một cấu trúc cố định.
### 2.3. Cập nhật và Quản lý document
- MongoDB cung cấp nhiều cách để quản lý dữ liệu trong document thông qua các thao tác CRUD (Create, Read, Update, Delete). Một số tính năng cụ thể:

- Insert (Thêm dữ liệu): Thêm tài liệu mới vào collection. MongoDB tự động sinh khóa chính _id cho mỗi document để đảm bảo duy nhất.

Ví dụ:
```
db.products.insert({
  name: "Laptop",
  brand: "Dell",
  price: 1200
})
```
- Update (Cập nhật dữ liệu): MongoDB cho phép cập nhật toàn bộ tài liệu hoặc chỉ cập nhật từng trường cụ thể. Điều này giúp cải thiện hiệu suất vì chỉ các trường cần thiết được thay đổi mà không cần ghi đè toàn bộ tài liệu.

Ví dụ: Cập nhật giá của sản phẩm:
``` 
db.products.update({ name: "Laptop" }, { $set: { price: 1100 } }) 
```
- Delete (Xóa dữ liệu): Cho phép xóa tài liệu dựa trên điều kiện.

Ví dụ: Xóa một sản phẩm dựa trên ID:
```
db.products.remove({ _id: ObjectId("60e8bc1") })

```
### 2.4. Indexing (Chỉ mục)
MongoDB cho phép tạo chỉ mục để tăng tốc độ truy vấn trên các trường cụ thể, cải thiện hiệu năng khi làm việc với các tập dữ liệu lớn.

Ví dụ:

```
db.products.createIndex({ name: 1 })

```

### 2.5. Replication và Sharding
- Replication (Nhân bản): MongoDB sử dụng replication để đảm bảo tính sẵn sàng và khôi phục dữ liệu. Một bản sao chính (primary) và nhiều bản sao phụ (secondary) được duy trì đồng bộ.
- Sharding (Phân mảnh dữ liệu): Dữ liệu lớn được chia thành nhiều phần và phân phối qua các shard, giúp hệ thống mở rộng và duy trì hiệu suất tốt hơn.
### 2.6. Aggregation (Tập hợp dữ liệu)
- MongoDB hỗ trợ Aggregation Framework để xử lý dữ liệu phức tạp, chẳng hạn như lọc, nhóm, và sắp xếp dữ liệu trực tiếp trong cơ sở dữ liệu mà không cần phải lấy tất cả dữ liệu ra bên ngoài.

Ví dụ:
```
db.sales.aggregate([
  { $match: { status: "completed" } },
  { $group: { _id: "$customer_id", total: { $sum: "$amount" } } }
])
```
- Cách quản lý dữ liệu trong MongoDB dựa trên tính linh hoạt của các document, cho phép lưu trữ và xử lý dữ liệu phi cấu trúc hiệu quả. Các công cụ mạnh mẽ như chỉ mục, phân mảnh và aggregation hỗ trợ quản lý dữ liệu với khối lượng lớn và tốc độ cao, phù hợp với các ứng dụng hiện đại.


## 3. Cách Collection tương tác với Document trong MongoDB
- Trong MongoDB, Collections và Documents tương tác với nhau theo mô hình cơ sở dữ liệu hướng tài liệu (Document-Oriented), giúp quản lý dữ liệu linh hoạt và tối ưu hóa cho các hệ thống phi cấu trúc.

### 3.1. Document trong Collection
- Document là đơn vị cơ bản lưu trữ dữ liệu trong MongoDB. Mỗi document chứa dữ liệu dưới dạng BSON (Binary JSON) và có thể bao gồm nhiều trường (fields) khác nhau với cấu trúc linh hoạt.
- Mỗi document trong MongoDB tương đương với một hàng trong cơ sở dữ liệu quan hệ, nhưng không cần tuân thủ một lược đồ cứng nhắc (schema-less), nghĩa là mỗi document trong cùng một collection có thể có cấu trúc khác nhau.
### 3.2. Collection lưu trữ Document
- Collection là tập hợp các document. Nó có chức năng tương đương với bảng (table) trong cơ sở dữ liệu quan hệ, nhưng các document trong collection không cần phải có cùng một cấu trúc hoặc kiểu dữ liệu.
Ví dụ: Trong một collection customers, một document có thể chỉ chứa tên và email của khách hàng, trong khi một document khác có thể bao gồm thêm địa chỉ và số điện thoại.
### 3.3. Tương tác giữa Collection và Document qua CRUD
- MongoDB cung cấp các thao tác CRUD (Create, Read, Update, Delete) để tương tác với document trong collection:

- Create (Thêm tài liệu mới):

- Document được thêm vào collection với một khoá chính mặc định là _id.
Ví dụ:
```
db.customers.insert({ name: "Nguyen dien sy dao", email: "24410013@ms.uit.edu.vn" })
```
- Read (Đọc dữ liệu):

- Truy vấn document từ collection dựa trên các tiêu chí cụ thể.
Ví dụ: Lấy tất cả khách hàng có tên là "Nguyen dien sy dao":
```
db.customers.find({ name: "Nguyen dien sy dao" })
```

- Update (Cập nhật dữ liệu):

- Cập nhật các trường cụ thể trong document mà không cần cập nhật toàn bộ tài liệu.
Ví dụ: Cập nhật email của một khách hàng:
```
db.customers.update({ name: "Nguyen dien sy dao" }, { $set: { email: "24410013@ms.uit.edu.vn" } })
```
- Delete (Xóa tài liệu):

- Xóa document khỏi collection dựa trên điều kiện nhất định.
Ví dụ: Xóa khách hàng có _id là "12345":

```
db.customers.remove({ _id: ObjectId("12345") })
```

### 3.4. Schema Validation trong Collection
- Mặc dù MongoDB không yêu cầu schema cố định, nhưng người dùng có thể thiết lập Schema Validation để áp dụng các quy tắc trên dữ liệu trong collection. Điều này đảm bảo tính toàn vẹn của dữ liệu.

Ví dụ: Xác thực rằng một document trong collection phải có trường "email" theo định dạng hợp lệ.
### 3.5. Indexes trong Collection
MongoDB cho phép tạo chỉ mục (index) trên các trường trong document để tối ưu hóa truy vấn và cải thiện hiệu năng.

Ví dụ:

```
db.customers.createIndex({ email: 1 })
```

### 3.6. Aggregation Framework
- *Aggregation Framework* cho phép thực hiện các phép tính phức tạp trên document trong collection mà không cần phải lấy tất cả dữ liệu ra khỏi cơ sở dữ liệu.

Ví dụ: Tính tổng số đơn hàng của mỗi khách hàng:

```
db.orders.aggregate([
  { $group: { _id: "$customerId", total: { $sum: "$amount" } } }
])
```

## 4. Lý do MongoDB sử dụng BSON
### 4.1. Hiệu suất cao và Tối ưu hóa bộ nhớ
- BSON (Binary JSON) được MongoDB lựa chọn làm định dạng lưu trữ chính không chỉ vì nó giúp tăng tốc độ xử lý dữ liệu mà còn do khả năng tối ưu hóa bộ nhớ của nó. BSON mã hóa dữ liệu dưới dạng nhị phân, giúp dữ liệu có thể được tuần tự hóa (serialization) nhanh chóng, dễ dàng và gọn gàng hơn so với JSON thuần túy.

- Một trong những vấn đề của JSON là tính thiếu tối ưu trong việc lưu trữ. JSON sử dụng cấu trúc chuỗi văn bản (text-based) để biểu diễn dữ liệu, điều này tạo ra nhiều dữ liệu không cần thiết khi lưu trữ hoặc truyền tải. Ví dụ: trong JSON, mỗi trường của tài liệu đều phải chứa tên thuộc tính bằng chuỗi, điều này dẫn đến việc tiêu tốn nhiều băng thông và dung lượng hơn.

- BSON khắc phục được vấn đề này bằng cách sử dụng dạng nhị phân để lưu trữ dữ liệu, giúp:

  - Giảm kích thước tệp khi truyền qua mạng.
  - Nén dữ liệu tốt hơn, từ đó tiết kiệm dung lượng đĩa.
  - Tăng tốc độ tuần tự hóa và giải tuần tự hóa dữ liệu (deserialization).
### 4.2. Hỗ trợ kiểu dữ liệu phong phú
- Một yếu tố khác biệt quan trọng của BSON so với JSON là hỗ trợ đa dạng kiểu dữ liệu. JSON cơ bản chỉ hỗ trợ các kiểu dữ liệu đơn giản như chuỗi (string), số (number), mảng (array), boolean, và null. Điều này tạo ra giới hạn khi làm việc với các kiểu dữ liệu phức tạp, ví dụ như:

  - Dữ liệu thời gian thực (timestamps): JSON không có kiểu dữ liệu ngày và giờ chuẩn, buộc các hệ thống phải xử lý bằng cách dùng chuỗi hoặc số nguyên, điều này không tối ưu cho truy vấn và xử lý.
  - Số lớn (64-bit integers): JSON bị giới hạn bởi chuẩn IEEE 754 nên không thể xử lý số nguyên lớn hơn 53 bit chính xác, trong khi BSON hỗ trợ Int64 và Decimal128, giúp làm việc với dữ liệu tài chính hoặc khoa học một cách hiệu quả.
- BSON bổ sung nhiều kiểu dữ liệu bổ sung như:

  - `Date`: Hỗ trợ đầy đủ kiểu ngày giờ.
  - `Binary data`: Hỗ trợ dữ liệu nhị phân, điều này rất cần thiết cho việc lưu trữ các tệp dữ liệu lớn như hình ảnh, file âm thanh hoặc mã hóa bảo mật.
  - `Embedded documents`: Khả năng lưu trữ các tài liệu lồng nhau (nested documents) mà không làm giảm hiệu năng.
Điều này giúp MongoDB có thể xử lý nhiều loại dữ liệu phức tạp hơn, từ đó đáp ứng được yêu cầu của các hệ thống hiện đại với khối lượng và kiểu dữ liệu đa dạng.

### 4.3. Khả năng phân tích và truy vấn nhanh chóng
- BSON được tối ưu hóa cho việc truy xuất một phần tài liệu (partial retrieval), giúp giảm thiểu tài nguyên khi xử lý dữ liệu. Điều này cực kỳ quan trọng trong môi trường cơ sở dữ liệu lớn, nơi MongoDB thường xuyên được sử dụng để lưu trữ hàng triệu đến hàng tỷ bản ghi.

- Trong một hệ thống cơ sở dữ liệu truyền thống, khi thực hiện truy vấn, toàn bộ tài liệu cần được tải vào bộ nhớ, điều này gây ra sự chậm trễ và tốn tài nguyên. Với BSON, MongoDB có thể chỉ cần tải phần dữ liệu được yêu cầu, nhờ vậy:

  - Giảm thiểu lượng tài nguyên xử lý cần thiết (RAM, CPU).
  - Tăng tốc độ truy vấn và cải thiện hiệu năng tổng thể.
- Hơn nữa, BSON có cấu trúc tự mô tả, tức là nó không chỉ lưu trữ dữ liệu mà còn lưu trữ loại dữ liệu của các trường. Điều này cho phép MongoDB truy vấn, sắp xếp, và lọc dữ liệu một cách hiệu quả mà không cần phải xử lý lại toàn bộ cấu trúc dữ liệu.

### 4.4. Tính mở rộng và linh hoạt
- Trong các hệ quản trị cơ sở dữ liệu quan hệ (RDBMS), dữ liệu được lưu trữ trong các bảng với cấu trúc cứng nhắc, yêu cầu xác định trước lược đồ dữ liệu (schema). Điều này gây khó khăn trong việc thay đổi hoặc mở rộng cấu trúc dữ liệu khi hệ thống phát triển.

- BSON, ngược lại, có tính linh hoạt cao vì không cần lược đồ cố định. Mỗi tài liệu trong MongoDB có thể có cấu trúc khác nhau, cho phép dễ dàng mở rộng khi thêm các thuộc tính mới mà không cần thay đổi cơ sở dữ liệu hiện có. Nhờ vào cấu trúc tự mô tả, mỗi trường trong tài liệu BSON đều bao gồm thông tin về kiểu dữ liệu của nó, giúp quản lý dữ liệu một cách tự động và dễ dàng.

- Điều này rất quan trọng đối với các ứng dụng web hoặc mobile hiện đại, nơi yêu cầu phát triển và thay đổi nhanh chóng, cần linh hoạt trong việc lưu trữ và quản lý dữ liệu.

### 4.5. Khả năng tương thích cao với JSON
- Cuối cùng, mặc dù BSON là định dạng nhị phân, nó vẫn giữ được sự dễ dàng trong việc thao tác với JSON. Do BSON là một mở rộng của JSON, các lập trình viên có thể dễ dàng sử dụng MongoDB mà không cần phải học thêm một ngôn ngữ mới. Điều này giúp tăng tốc độ phát triển, đồng thời vẫn đảm bảo hiệu quả khi lưu trữ và truyền dữ liệu.

- JSON vẫn là định dạng phổ biến trong các API web và các ứng dụng khác, và việc BSON tương thích với JSON giúp MongoDB dễ dàng tích hợp với các hệ thống hiện có.

**Kết luận:**
- MongoDB lựa chọn BSON vì khả năng tối ưu dung lượng, tốc độ xử lý nhờ mã hóa nhị phân, và hỗ trợ nhiều kiểu dữ liệu phức tạp mà JSON không đáp ứng được. BSON cải thiện hiệu suất truy vấn và giảm thiểu tài nguyên, rất quan trọng trong các hệ thống xử lý khối lượng dữ liệu lớn. Đồng thời, tính tự mô tả của BSON giúp MongoDB dễ dàng mở rộng và thích ứng với thay đổi mà không cần tái cấu trúc dữ liệu. Khả năng tương thích với JSON giúp BSON trở thành lựa chọn mạnh mẽ và thân thiện với nhà phát triển.


## Nguồn tham khảo
- [MongoDB Manual: BSON Types](https://www.mongodb.com/docs/manual/reference/bson-types/)
- [MongoDB Replica Set](https://www.mongodb.com/docs/manual/replication/)
- [MongoDB Sharding](https://www.mongodb.com/docs/manual/sharding/)
- [MongoDB Customer Stories](https://www.mongodb.com/solutions/customer-case-studies)

- [MongoDB Official Website](https://www.mongodb.com/).
- [MongoDB Wiki trên GitHub](https://github.com/mongodb/mongo/wiki).
- [MongoDB Blog](https://www.mongodb.com/blog).
- [MongoDB Documentation](https://www.mongodb.com/docs/).
- [MongoDB Annual Report 2024](https://investors.mongodb.com/2024-Annual-Report-and-Proxy-Statement).
- [Stock Analysis](https://stockanalysis.com/stocks/mdb/revenue/).
- [MongoDB Inc. Performance Reports](https://investors.mongodb.com/news-releases/news-release-details/mongodb-inc-announces-first-quarter-fiscal-2025-financial).
- [Stack Overflow Developer Survey 2024](https://survey.stackoverflow.co/2024/technology).
- [Technology | 2024 Stack Overflow Developer Survey](https://survey.stackoverflow.co/2024/technology#most-popular-technologies-database-prof).
- [DB-Engines Ranking](https://db-engines.com/en/ranking).
