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
  "name": "Alice",
  "age": 28,
  "email": "alice@example.com"
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
  "name": "Alice",
  "age": 30,
  "isStudent": false,
  "courses": ["Mathematics", "Computer Science"],
  "address": {
    "street": "123 Elm Street",
    "city": "Springfield",
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
+ Courses: ["Mathematics", "Computer Science"] (dưới dạng mảng chuỗi nhị phân).
+ Address: {
+ Street: "123 Elm Street" (dưới dạng chuỗi nhị phân).
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


<!-- Phần chưa thực hiện  -->
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

## Nguồn tham khảo

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
