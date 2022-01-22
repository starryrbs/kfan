<template>
  <div>
    <ul>
      <li v-for="house in houses">
        <div class="clearfix" style="height: 235px;padding: 20px">
          <div class="nlc_img">
            <a data-showlocation="2" data-tabtype="新房" target="_blank"
               href="http://sh.newhouse.fang.com/loupan/1210185694.htm">
              <img class="img_h" alt="中国铁建·花语江南"
                   :src="house.image">
            </a>
            <div class="inf_con">

              <span class="img_xgt"></span>


              <a class="icon_video"></a>


              <a class="icon_VR"></a>

            </div>
          </div>
          <div class="nlc_details">
            <div class="house_price">

              <span>价格待定</span>

            </div>
            <div class="nlcd_name">
              <a data-showlocation="2" data-tabtype="新房" target="_blank"><span class="tit_shop" style="cursor: pointer"
                                                                               @click="()=>{onClickHouseTitle(house)}">{{
                  house.title
                }}</span></a>
            </div>
            <div class="house_type clearfix">


              <a>{{ house.description }}</a>
              &nbsp;
              <span>4.14分/20条评论</span>
            </div>
            <div class="address">
              <span class="sngrey">青浦-赵巷</span>
              <i></i>业文路129弄
            </div>
            <div class="stag_box clearfix">

              <span class="forSale">待售</span>


              <span>品牌地产</span>


            </div>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>

import {defineComponent} from 'vue'
import {getHouses} from "@/api/house";
import {GetUser, StoreUser} from "@/utils/user";
import {saveHistory} from "@/api/history";

export default defineComponent({
  name: "index",
  data() {
    return {
      houses: []
    }
  },
  mounted() {
    getHouses().then(response => {
      this.houses = response.data.results
    })
  },
  methods: {
    onClickHouseTitle(house) {
      const user = GetUser()
      if (user) {
        saveHistory(user.id, house.id, "house")
      }
    }
  }
})
</script>

<style scoped lang="less">
.nlc_img {
  position: relative;
  float: left;
  margin-right: 26px;
  width: 180px;
  height: 135px;
  border-radius: 5px;
}

.nlc_details {
  float: left;
  position: relative;
  padding-top: 1px;
  width: 544px;
  min-height: 135px;

  .house_price {
    float: right;
    margin-top: 50px;
    line-height: 29px;
    color: #df2f30;
  }

  .nlcd_name {
    margin-top: -4px;
    height: 30px;
    line-height: 30px;
    font-size: 20px;
    font-weight: 700;
    color: #1c1b1b;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .house_type {
    position: relative;
    margin-top: 8px;
    width: 344px;
    height: 30px;
    line-height: 30px;
    font-size: 14px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .address {
    margin-top: 8px;
    width: 344px;
    line-height: 21px;
    color: #999;
    font-size: 14px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .stag_box {
    margin-top: 18px;

    .forSale {
      color: #f59505;
      background: #fff6e8;
    }

    span {
      display: block;
      float: left;
      margin-right: 8px;
      padding: 0 8px;
      height: 24px;
      line-height: 24px;
      color: #666;
      background: #f4f5f6;
    }
  }
}


</style>