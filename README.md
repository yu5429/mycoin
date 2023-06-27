# mycoin project
블록체인 구현 완료 <br>
노마드 코더 강의를 통해서 Go로 제작한 블록체인입니다. <br>
같이 제작된 rest api로 블록체인을 추가, 확인 할 수 있습니다. <br>

<br><br>
코인도 차후 구현 예정
<br><br><br><br><br>
**블록체인 구현 모습**<br><br>

***
main 실행<br>
![image](https://github.com/yu5429/mycoin/assets/123722364/2dfd3b38-2a4b-409f-987d-ad485f327607)
<br><br><br><br>
***

**각 URL별 기능**
4000: 기본 화면<br>
4000/blocks :전체 블록 보기<br>
(POST)4000/blocks /새로운 블록 추가<br>
4000/bloks/+찾고 싶은 블록 Hash :해당블록 확인<br>
![image](https://github.com/yu5429/mycoin/assets/123722364/e2f91f04-3d4f-4d83-9f24-13e8afa12e4a)
<br>
***
**전체 블록 확인**<br>
최초 확인 시 기본 블록 "Genesis"만 확인<br>
![image](https://github.com/yu5429/mycoin/assets/123722364/544b4309-8280-424d-a848-b2153b2bdac7)
<br>
***
**새로운 블록 추가**<br>
(POST)4000/blocks URL request로 새로운 블록 추가<br>
![image](https://github.com/yu5429/mycoin/assets/123722364/abf7962d-f8f6-44df-95ed-cd0260c6f701)
<br>
***
**추가된 블록 확인**<br>
4000/blocks URL request로 추가된 블록 확인<br>
![image](https://github.com/yu5429/mycoin/assets/123722364/959b1393-b66d-496e-91f5-1fcd599944f7)
<br><br>
**계속 추가된  블록**<br>
![image](https://github.com/yu5429/mycoin/assets/123722364/84f2f509-c724-406c-a49e-fb15a3bc01ba)
<br>
***
**특정 블록 확인**<br>
4000/bloks/+찾고 싶은 블록 Hash값 입력 후 Send Requst<br>
![image](https://github.com/yu5429/mycoin/assets/123722364/bac17f45-076f-44cd-8cd9-cc7989d8f623)

