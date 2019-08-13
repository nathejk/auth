<template>
    <main>
    <div class="box">
        <h1>Nathejk 2019</h1>
        <form :action="action" method="post" >
            <div class="input-user"><input type="text" name="username" placeholder="brugernavn"></div>
            <div class="input-pass"><input type="password" name="password" placeholder="kodeord"></div>
            <button>luk mig ind &raquo;</button>
        </form>
    </div>
    </main>
</template>

<style lang="scss">
@-webkit-keyframes autofill {
    to {
        color: #666;
        background: transparent;
    }
}
.box {
    height:300px;
    width:600px;
    padding:5px 15px;
    background-color:rgba(0,0,0,.8);
    -webkit-border-radius: 10px;
    -moz-border-radius: 10px;
    border-radius: 10px;

    font-family: 'Context Reprise Condensed SSi', sans-serif;

    h1 {
        background: url('/moon620.png') no-repeat;
        background-size: 80px 80px;
        color:#cc0;
        font-family: 'Context Reprise Condensed SSi';
        padding:35px 0px 0px 70px;
    }
    form {
        padding-left:250px;
        font-size:1.3rem;

        .input-user, .input-pass {
            background-color:#999;
            background-repeat-x: no-repeat;
            background-repeat-y: no-repeat;
            background-size: 40px 40px;
            background-position-y: center;
            padding:3px 0px 3px 50px;
            margin:10px 0px;
        }
        .input-user {
            background-image: url('/icons/glyphicons-basic-4-user@2x.png');
        }
        .input-pass {
            background-image: url('/icons/glyphicons-basic-217-lock@2x.png');
        }
        input {
            width:100%;
            background:transparent;
            border:0;

            &:focus {
                outline: none;
            }
            &:-webkit-autofill,
            &:-webkit-autofill:hover,
            &:-webkit-autofill:focus,
            &:-webkit-autofill:active {
                -webkit-animation-name: autofill;
                -webkit-animation-fill-mode: both;
                -webkit-text-fill-color: black !important;
            }
        }
        button {
            width: 100%;
            color: #000;
            background-color: #999;
            border: none;
            border-radius: 5px;
            box-shadow: 0px 8px 15px rgba(0, 0, 0, 0.1);
            transition: all 0.3s ease 0s;
            outline: none;
            padding: 3px 0px;
            margin:30px 10px 10px 0px;

            &:hover, &:focus {
              background-color: #cc0;
              box-shadow: 0px 15px 20px rgba(204, 204, 0, 0.4);
              transform: translateY(-7px);
            }
        }
    }
}
main {
    display:flex;
    justify-content: center;
    align-items: center;
    height:100%;
}
</style>

<script>
//  import '@/assets/main.scss'
import Card from '@/components/Card.vue'
import moment from 'moment'
import axios from 'axios'

export default {
    data: () => ({
      title: 'Nathejk 2019',
      countdownDate: moment('2019-04-26 22:32'),
      now: moment(),
      tick: null,
      user: {
          username: '',
          password: '',
      },
    }),
    components: {  Card },
    computed: {
        countdownActive: function() {
            //return this.countdownDate.isBefore(moment());
            return moment().isBefore(this.countdownDate);
        },
        action: function() {
            return envConfig.API_BASEURL + "/basicauth";
        },
    },
    methods: {
      afterEnter: function () {
        // Scroll to top on page switch
        window.scrollTo({top: 0, behavior: 'smooth'})
      },
      signin: function() {

      },
      submit: async function(e) {
        e.preventDefault();
        try {
            const rsp = await axios.post('//api.auth.local.nathejk.dk/basicauth', this.user)
        } catch(error) {
            throw new Error(error.response.data)
        }
        console.log(this.user);
      },
    },
    mounted: function () {
        this.tick = setInterval(function () {
            this.now = moment();
        }.bind(this), 1000)
    },
    beforeDestroy() {
        clearInterval(this.tick);
    }
}
</script>
