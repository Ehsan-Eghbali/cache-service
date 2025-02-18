
# میکروسرویس کش در حافظه (In-Memory Cache) با Go

در این پروژه، یک میکروسرویس ساده با زبان **Go** پیاده‌سازی شده که وظیفه‌ی ذخیره و بازیابی داده‌ها در حافظه را بر عهده دارد. این کار باعث بهبود سرعت پاسخ‌دهی و کاهش بار روی دیتابیس یا سایر سرویس‌های پشتیبان می‌شود.

## فهرست مطالب

1. [ویژگی‌های پروژه](#ویژگیهای-پروژه)
2. [پیش‌نیازها](#پیشنیازها)
3. [نصب و راه‌اندازی](#نصب-و-راهاندازی)
4. [ساختار پروژه](#ساختار-پروژه)
5. [استفاده](#استفاده)
6. [کتابخانه‌های جایگزین](#کتابخانههای-جایگزین)
7. [لایسنس](#لایسنس)

---

## ویژگی‌های پروژه

- **سرعت بالا در پاسخ‌دهی:** نگهداری داده‌های پرتکرار در حافظه منجر به کاهش زمان واکشی می‌شود.
- **مقیاس‌پذیری مناسب:** با کاستن از درخواست‌های مستقیم به دیتابیس و سرویس‌های دیگر، می‌توان در مصرف منابع صرفه‌جویی کرد.
- **پیاده‌سازی ساده:** از کتابخانه‌ی [go-cache](https://github.com/patrickmn/go-cache) برای مدیریت کش و زمان انقضا استفاده شده است.
- **ارائه سرویس مستقل:** به‌صورت یک میکروسرویس مجزا عمل می‌کند و می‌تواند در کنار سایر سرویس‌ها فعالیت کند.

---

## پیش‌نیازها

- نصب [Go](https://golang.org/doc/install) نسخه 1.18 یا بالاتر
- دسترسی به **Git** برای کلون یا دریافت سورس‌کد پروژه (در صورت نیاز)

---

## نصب و راه‌اندازی

1. مخزن پروژه را کلون کنید یا فایل‌ها را دانلود کنید:

   ```bash
   git clone https://github.com/Ehsan-Eghbali/cache-service.git
   ```

2. وارد دایرکتوری پروژه شوید:

   ```bash
   cd cache-service
   ```

3. وابستگی‌های پروژه را نصب کنید:

   ```bash
   go mod tidy
   ```

4. سرویس را اجرا کنید:

   ```bash
   go run main.go
   ```

5. در نهایت، سرویس روی پورت پیش‌فرض `8080` در حال اجرا خواهد بود:

   ```
   Starting in-memory cache microservice on port 8080...
   ```

---

## ساختار پروژه

```
in-memory-cache-microservice/
│
├── main.go          # فایل اصلی شامل کد راه‌اندازی سرویس
├── go.mod           # فایل مربوط به ماژول Go
├── go.sum           # فایل جمع‌آوری وابستگی‌های دقیق (hash و نسخه)
└── README.md        # فایل توضیحات و مستندات پروژه
```

---

## استفاده

پس از راه‌اندازی سرویس، می‌توانید از طریق متدهای HTTP به آن درخواست ارسال کنید:

- **آدرس پیش‌فرض:** `http://localhost:8080/data`
- **پارامتر ورودی (GET):**
    - `key`: کلید داده‌ای که قصد واکشی آن را دارید.

نمونه درخواست:

```bash
curl "http://localhost:8080/data?key=myKey"
```

### سناریوی پاسخ:

1. اگر **کلید** در کش وجود داشته باشد:

   ```json
   {
     "key": "myKey",
     "value": "value for myKey",
     "source": "cache"
   }
   ```
   مقدار `source` نشان می‌دهد که داده از **کش** بازیابی شده است.

2. اگر **کلید** در کش وجود نداشته باشد:
    - ابتدا داده از منبع اصلی (مثلاً دیتابیس) واکشی می‌شود (شبیه‌سازی شده با تابع `fetchData`).
    - سپس در کش ذخیره و به کلاینت بازگردانده می‌شود:

   ```json
   {
     "key": "myKey",
     "value": "value for myKey",
     "source": "database"
   }
   ```
   مقدار `source` نشان می‌دهد که داده از **منبع اصلی** بازیابی شده است.

---

## کتابخانه‌های جایگزین

علاوه بر `go-cache`، کتابخانه‌های دیگری هم برای مدیریت In-Memory Cache وجود دارند که می‌توانید بسته به نیازهای پروژه از آن‌ها استفاده کنید:

1. **[BigCache](https://github.com/allegro/bigcache):** برای حجم بالای داده و سرعت خواندن/نوشتن بالا بهینه شده است.
2. **[FreeCache](https://github.com/coocood/freecache):** با ساختاری بهینه برای موازی‌سازی (Concurrency) بالا و جلوگیری از تداخل گوروتین‌ها.
3. **[Groupcache](https://github.com/golang/groupcache):** مناسب برای کش توزیع‌شده (Distributed Cache)، توسط گوگل توسعه داده شده است.
4. **[Ristretto](https://github.com/dgraph-io/ristretto):** رویکردی هوشمندانه برای کنترل اندازه کش و کاهش هدررفت حافظه.

---

## لایسنس

این پروژه تحت لایسنس **MIT** ارائه می‌شود. برای اطلاعات بیشتر، به فایل [LICENSE](LICENSE) (در صورت وجود) یا اطلاعات لایسنس در انتهای سورس‌کد مراجعه کنید.

---

> **نکته:** لطفاً با توجه به نیازمندی‌ها و سیاست‌های خود، زمان انقضا، مکانیزم پاکسازی کش و ساختار لاگینگ را سفارشی‌سازی کنید. همچنین در محیط‌های واقعی تولید (Production) ممکن است نیاز باشد از استراتژی‌های توزیع‌شده یا ترکیبی (Redis و غیره) برای مدیریت بهتر کش بهره ببرید.
