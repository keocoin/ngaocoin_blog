---
Title: "Tạo địa chỉ ví an toàn cho hodler bằng smartphone cũ"
Category: Cơ bản
Slug: tao-dia-chi-vi-hodler
Thumb: https://i.ibb.co/5RQ394f/private.png
Author: Tào Mạnh Đức
PublishDate: "2023-02-11"
---

Như anh em đã biết ví lạnh là ví an toàn nhất để lưu trữ tiền mã hóa (coin). Tuy nhiên câu hỏi là làm sao để tạo ra ví lạnh? và làm sao để lưu trữ ví lạnh đúng cách?

Câu trả lời dễ nhất là mua ví lạnh từ các nhà cung cấp tên tuổi. Hiện tai nếu bạn ở Vietnam thì hình như chưa có nhà cung cấp chính thức, chủ yếu là mua qua bên thứ 3 hoặc xách tay về.

Hôm nay mình sẽ share cho bạn một cách dễ dàng hơn để tận dụng điện thoại smartphone cũ (không còn dùng nữa) để tạo ra địa chỉ ví phục vụ cho việc hodl coin lâu dài

## Nguyên tắc hoạt động cơ bản của Ví

Về cơ bản, 1 ví có 2 thành phần chính: Khóa cá nhân, và địa chỉ ví.

### Khóa cá nhân là gì?

Nếu bạn dùng các ví như Trust hoặc Metamask, cụm từ hạt giống không phải là khóa cá nhân, nó chỉ là một ánh xạ của khóa cá nhân.

Về mặt kỹ thuật, khóa cá nhân là một chuỗi 64 ký tự, được tạo thành bằng việc random một số từ 1 đến 2^256. Sau đó dùng số này đưa vào hàm sha256.

Ví dụ: nếu random ra số ```1```, thì khóa cá nhân sẽ là ```6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b```, còn random ra số ```2``` thì khóa cá nhân sẽ là ```d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35```

Từ khóa cá nhân này, phần mềm ví sẽ tiếp tục mã hóa thêm vài bược nữa để ra một địa chỉ cuối cùng. Chính là địa chỉ mà chúng ta dùng hàng ngày để nhận tiền.

Khi bạn chuyển tiền, thì phần mềm ví sẽ dùng khóa cá nhân này để tạo ra chữ ký cho giao dịch.

Vậy nên ai có khóa cá nhân đều có thể tạo ra chữ ký và rút tiền trong ví của bạn.

### Vậy thì cụm từ hạt giống là từ đâu ra?

Vì nhớ khóa cá nhân rất khó, và dễ sai, nên người ta nghĩ ra một cách để biến khá cá nhân thành các chuỗi ký tự mà con người dễ đọc, và có thể nhớ được (cụm từ hạt giống). Đồng thời từ ```cụm từ hạt giống```, chúng ta có thể decode ngược lại để lấy lại khá cá nhân.

Từ đó chuẩn BIP ra đời, hiện tại có nhiều chuẩn BIP, nhưng chủ yếu sẽ là BIP39, hiện được các bên phát triển ví sử dụng.

### Thông tin thêm

Trên thực tế, chúng ta thường không dùng cách này để tọa ra ví cá nhân, tuy nhiên, để cho bạn nào có ý định nghịch thử, mình sẽ dẫn link bên dưới.

Một số trang web các bạn có thể dùng để thử:


[https://emn178.github.io/online-tools/sha256.html ](https://emn178.github.io/online-tools/sha256.html ) Tạo một khóa cá nhân từ chữ số hoặc ký tự.
[https://privatekeys.pw/calc](https://privatekeys.pw/calc) <= Lấy địa chỉ ví từ khóa cá nhân.
[https://iancoleman.io/bip39/#english](https://iancoleman.io/bip39/#english) <= Khóa cá nhân thành cụm từ hạt giống.

## Tận dụng điện thoại cũ để tạo ra ví an toàn thế nào?

Như các bạn thấy, nếu ai đó biết được cụm từ hạt giống hoặc khóa cá nhân, thì họ có thể dùng nó để login vào metamask hoặc trustwallet và rút tiền của bạn. Hoặc máy tính của bạn bị dính virrus, và bạn bị lộ khá cá nhân.

Làm thế nào để tạo ra khóa cá nhân an toàn?

Có một cách đơn giản là lúc tạo ví, bạn tạo offline, ghi ra giấy sau đó xóa mọi thông tin trên máy liên quan đến ví đó đi. Cách này thì vẫn còn một rủi ro là máy tính bạn đã dính virus từ trước, khi bạn online trở lại, con virus này sẽ gửi các thông tin đã lấy đc về cho chủ nó.

### Tạo ví an toàn từ smartphone cũ

- reset toàn bộ máy, sau đó cài ví trustwallet.
- Chuyển qua chế độ máy bay (offline)
- Tạo ví mới, và ghi cụm từ hạt giống ra giấy hoặc ví cứng bằng sắt (các bạn nên lưu làm 2 3 phiên bản khác nhau, ở nhiều nơi an toàn khác nhau).
- Dùng một điện thoại khác chụp lại địa chỉ ví (chỉ chụp lại địa chỉ ví chứ không phải là cụm từ hạt giống)
- Reset máy một lần nữa

Bằng cách này, bạn có thể yên tâm là cụm từ hạt giống của bạn hoàn toàn offline. Cách này phù hợp cho mục đích hodl coin lâu dài, khoảng 1 2 năm trở lên.

Khi bạn muốn nhận tiền thì dùng địa chỉ ví đã chụp lại để gửi tiền vào. Khi bạn muốn chuyển tiền ra khỏi ví thì dùng cụm từ hạt giống để khôi phục ví. Và để an toàn thì nên dùng 1 lần, do khi đã login lại bằng điện thoại đã kết nối thì ví này không còn là ví lạnh nữa.

## Kết luận

Cách làm này khá bất tiện, tuy nhiên bạn tiết kiệm đc việc mua một chiếc ví lạnh, và rất an toàn, do không kết nối với internet.

Nếu bạn có cách nào hay hơn, vui lòng comment bên dưới cho anh em khác học theo.