import Vue from 'vue';
import VueRouter from 'vue-router';
import login from '@/views/login/index';
import admin from '@/views/admin/index';
import adminHome from '@/views/admin/home';
import studentManage from '@/views/admin/studentManage/index'
import addStudent from "@/views/admin/studentManage/addStudent";
import studentList from "@/views/admin/studentManage/studentList";
import editorStudent from "@/views/admin/studentManage/editorStudent";
import teacherManage from "@/views/admin/teacherManage/index"
import addTeacher from "@/views/admin/teacherManage/addTeacher";
import editorTeacher from "@/views/admin/teacherManage/editorTeacher";
import courseManage from "@/views/admin/courseManage/index";
import addCourse from "@/views/admin/courseManage/addCourse";
import teacher from "@/views/teacher/index";
import queryStudent from "@/views/admin/studentManage/queryStudent";
import queryTeacher from "@/views/admin/teacherManage/queryTeacher";
import student from "@/views/student/index";
import editorCourse from "@/views/admin/courseManage/editorCourse";
import courseList from "@/views/admin/courseManage/courseList";
import queryCourse from "@/views/admin/courseManage/queryCourse";
import offerCourse from "@/views/teacher/offerCourse";
import teacherHome from "@/views/teacher/home";
import setCourse from "@/views/teacher/setCourse";
import studentHome from "@/views/student/home";
import myOfferCourse from "@/views/teacher/myOfferCourse";
import CourseTeacherManage from "@/views/admin/selectCourseManage/index";
import queryCourseTeacher from "@/views/admin/selectCourseManage/queryCourseTeacher";
import studentSelectCourseManage from "@/views/student/selectCourse/index";
import selectCourse from "@/views/student/selectCourse/selectCourse";
import querySelectedCourse from "@/views/student/selectCourse/querySelectedCourse";
import studentCourseGrade from "@/views/student/courseGrade/index";
import queryCourseGrade from "@/views/student/courseGrade/queryCourseGrade";
import queryGradeCourse from "@/views/admin/gradeCourseManage/queryGradeCourse";
import editorGradeCourse from "@/views/admin/gradeCourseManage/editorGradeCourse";
import teacherGradeCourseManage from "@/views/teacher/teacherGradeCourseManage/index";
import teacherQueryGradeCourse from "@/views/teacher/teacherGradeCourseManage/teacherQueryGradeCourse";
import teacherEditorGradeCourse from "@/views/teacher/teacherGradeCourseManage/teacherEditorGradeCourse";
import updateInfo from "@/components/updateInfo";

Vue.use(VueRouter)

const routes = [
  {
    // ?????????????????????
    path: '/',
    name: 'index',
    component: login,
    redirect: '/login'
  },
  {
    // ?????????
    path: '/login',
    name: 'login',
    component: login
  },
  {
    // admin ?????????
    path: '/admin',
    name: 'admin',
    redirect: '/adminHome',
    component: admin,
    meta: {requireAuth: true},
    children: [
      {
        path: '/adminHome',
        name: 'Hi! admin',
        component: adminHome,
        meta: {requireAuth: true},
        children: [
          {
            path: '/adminHome',
            name: 'admin ??????',
            component: adminHome,
            meta: {requireAuth: true},
          }
        ]
      },
      {
        path: '/studentManage',
        name: '????????????',
        component: studentManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/addStudent',
            name: '????????????',
            component: addStudent,
            meta: {requireAuth: true}
          },
          {
            path: '/editorStudent',
            name: '????????????',
            component: editorStudent,
            meta: {requireAuth: true}
          },
          {
            path: '/queryStudent',
            name: '????????????',
            component: queryStudent,
            meta: {requireAuth: true},
            children: [
            ]
          }
        ]
      },
      {
        path: '/teacherManage',
        name: '????????????',
        component: teacherManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/addTeacher',
            name: '????????????',
            component: addTeacher,
            meta: {requireAuth: true}
          },
          {
            path: '/queryTeacher',
            name: '????????????',
            component: queryTeacher,
            meta: {requireAuth: true},
            children: [
            ]
          },
          {
            path: '/editorTeacher',
            name: '????????????',
            component: editorTeacher,
            meta: {requireAuth: true}
          },
        ]
      },
      {
        path: '/courseManage',
        name: '????????????',
        component: courseManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/addCourse',
            name: '????????????',
            component: addCourse,
            meta: {requireAuth: true}
          },
          {
            path: '/queryCourse',
            name: '????????????',
            component: queryCourse,
            meta: {requireAuth: true},
            children: [
              {
                path: '/courseList',
                name: '????????????',
                component: courseList,
                meta: {requireAuth: true}
              },
            ]
          },
          {
            path: '/editorCourse',
            name: '????????????',
            component: editorCourse,
            meta: {requireAuth: true}
          },
        ]
      },
      {
        path: '/CourseTeacher',
        name: '???????????????',
        component: CourseTeacherManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/queryCourseTeacher',
            name: '????????????',
            component: queryCourseTeacher,
            meta: {requireAuth: true},
          }
        ]
      },
      {
        name: 'admin ??????????????????',
        path: "/gradeCourseManage",
        component: studentManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/queryGradeCourse',
            name: '??????????????????',
            component: queryGradeCourse,
            meta: {requireAuth: true},
          },
          {
            path: '/editorGradeCourse',
            name: '??????',
            component: editorGradeCourse,
            meta: {requireAuth: true}
          }
        ]
      }
    ]
  },
  {
    path: '/teacher',
    name: 'teacher',
    component: teacher,
    redirect: '/teacherHome',
    meta: {requireAuth: true},
    children: [
      {
        path: '/teacherHome',
        name: 'Hi! teacher',
        meta: {requireAuth: true},
        component: teacherHome,
        children: [
          {
            path: '/teacherHome',
            name: '????????????',
            meta: {requireAuth: true},
            component: teacherHome
          },
        ]
      },
      {
        path: '/updateInfo',
        name: '????????????',
        component: updateInfo,
        meta: {requireAuth: true},
        children: [
          {
            path: '/updateInfoHome',
            name: '??????????????????',
            component: updateInfo,
            meta: {requireAuth: true}
          }
        ]
      },
      {
        path: '/courseManage',
        name: '????????????',
        meta: {requireAuth: true},
        component: setCourse,
        children: [
          {
            path: '/myOfferCourse',
            name: '??????????????????',
            component: myOfferCourse,
            meta: {requireAuth: true}
          },
          {
            path: '/offerCourse',
            name: '????????????',
            component: offerCourse,
            meta: {requireAuth: true}
          },
        ]
      },
      {
        name: '??????????????????',
        path: '/teacherQueryGradeCourseManage',
        component: teacherGradeCourseManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/teacherQueryGradeCourseManage',
            name: '????????????',
            component: teacherQueryGradeCourse,
            meta: {requireAuth: true}
          },
          {
            path: '/teacherEditorGradeCourse',
            name: '????????????',
            component: teacherEditorGradeCourse,
            meta: {requireAuth: true}
          }
        ]
      }
    ]
  },
  {
    path: '/student',
    name: 'student',
    component: student,
    redirect: '/studentHome',
    meta: {requireAuth: true},
    children: [
      {
        path: '/student',
        name: 'hi! student',
        component: studentHome,
        meta: {requireAuth: true},
        children: [
          {
            path: '/studentHome',
            name: '????????????',
            component: studentHome,
            meta: {requireAuth: true},
          },
        ],
      },
      {
        path: '/updateInfo',
        name: '????????????',
        component: updateInfo,
        meta: {requireAuth: true},
        children: [
          {
            path: '/updateInfoHome',
            name: '??????????????????',
            component: updateInfo,
            meta: {requireAuth: true}
          }
        ]
      },
      {
        path: '/studentSelectCourseManage',
        name: '????????????',
        component: studentSelectCourseManage,
        meta: {requireAuth: true},
        children: [
          {
            path: '/studentSelectCourse',
            name: '??????',
            component: selectCourse,
            meta: {requireAuth: true}
          },
          {
            path: '/querySelectedCourse',
            name: '????????????',
            component: querySelectedCourse,
            meta: {requireAuth: true}
          }
        ]
      },
      {
        path: '/courseGrade',
        name: '??????????????????',
        component: studentCourseGrade,
        meta: {requireAuth: true},
        children: [
          {
            path: '/queryCourseGrade',
            name: '????????????',
            component: queryCourseGrade,
            meta: {requireAuth: true}
          },
        ]
      }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

/*
  session ?????????
    1. token
    2. name
    3. type
    4. tid
    5. sid
    5. ???????????? info
 */
router.beforeEach((to, from, next) => {
  console.log(from.path + ' ====> ' + to.path)
  if (to.meta.requireAuth) { // ???????????????????????????????????????
    if (sessionStorage.getItem("token") === 'true') { // ????????????????????????token
      next()
    } else {
      // ?????????,?????????????????????
      next({
        path: '/login',
        query: {redirect: to.fullPath}
      })
    }
  } else {
    // ????????????????????????????????????????????????????????????????????????
    if(sessionStorage.getItem("token") === 'true'){
      console.log('???????????????????????????????????????????????????')
      const t = sessionStorage.getItem("type")
      next('/' + t);
    }else{
      next();
    }
  }
});